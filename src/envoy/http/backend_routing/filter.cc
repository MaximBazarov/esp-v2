// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// TODO(jcwang): Add tests
#include "src/envoy/http/backend_routing/filter.h"

#include <string>

#include "common/http/headers.h"
#include "src/envoy/utils/filter_state_utils.h"

namespace Envoy {
namespace Extensions {
namespace HttpFilters {
namespace BackendRouting {

using ::Envoy::StreamInfo::FilterState;
using ::google::api::envoy::http::backend_routing::BackendRoutingRule;
using Http::FilterDataStatus;
using Http::FilterHeadersStatus;
using Http::FilterTrailersStatus;
using Http::HeaderMap;
using Http::LowerCaseString;

Filter::Filter(FilterConfigSharedPtr config) : config_(config) {
  for (const auto& rule : config_->config().rules()) {
    backend_routing_map_[rule.operation()] = rule;
  }
}

FilterHeadersStatus Filter::decodeHeaders(HeaderMap& headers, bool) {
  const FilterState& filter_state =
      decoder_callbacks_->streamInfo().filterState();
  absl::string_view operation =
      Utils::getStringFilterState(filter_state, Utils::kOperation);
  // NOTE: this shouldn't happen in practice because Path Matcher filter would
  // have already rejected the request.
  if (operation.empty()) {
    ENVOY_LOG(debug, "No operation found from DynamicMetadata");
    return FilterHeadersStatus::Continue;
  }

  ENVOY_LOG(debug, "Found operation: {}", operation);
  auto it = backend_routing_map_.find(operation);
  if (it == backend_routing_map_.end()) {
    ENVOY_LOG(debug, "No backend routing rule found for operation {}",
              operation);
    return FilterHeadersStatus::Continue;
  }
  const BackendRoutingRule& rule = it->second;
  std::string newPath;
  ENVOY_LOG(debug, "backend routing for operation {}, old path: {}", operation,
            headers.Path()->value().getStringView());
  if (rule.is_const_address()) {  // CONSTANT_ADDRESS
    absl::string_view queryParamFromPathParam =
        Utils::getStringFilterState(filter_state, Utils::kQueryParams);
    const auto originalPath =
        std::string(headers.Path()->value().getStringView());
    std::size_t originalQueryParamPos = originalPath.find('?');
    if (originalQueryParamPos != std::string::npos) {
      // has query parameters in original url
      const std::string& originalQueryParam =
          originalPath.substr(originalQueryParamPos);
      newPath = absl::StrCat(rule.path_prefix(), originalQueryParam);
      if (!queryParamFromPathParam.empty()) {
        absl::StrAppend(&newPath, "&", queryParamFromPathParam);
      }
    } else {
      newPath = rule.path_prefix();
      if (!queryParamFromPathParam.empty()) {
        absl::StrAppend(&newPath, "?", queryParamFromPathParam);
      }
    }
    config_->stats().constant_address_request_.inc();
    ENVOY_LOG(debug,
              "constant address backend routing for operation {}, new path: {}",
              operation, newPath);
  } else {  // APPEND_PATH_TO_ADDRESS
    newPath = absl::StrCat(rule.path_prefix(),
                           headers.Path()->value().getStringView());
    config_->stats().append_path_to_address_request_.inc();
    ENVOY_LOG(
        debug,
        "append path to address backend routing for operation {}, new path: {}",
        operation, newPath);
  }
  const auto& pathField = Http::Headers::get().Path;
  headers.remove(pathField);
  headers.addCopy(pathField, newPath);

  return FilterHeadersStatus::Continue;
}

FilterDataStatus Filter::decodeData(Buffer::Instance&, bool) {
  return FilterDataStatus::Continue;
}

FilterTrailersStatus Filter::decodeTrailers(HeaderMap&) {
  return FilterTrailersStatus::Continue;
}

void Filter::setDecoderFilterCallbacks(
    Http::StreamDecoderFilterCallbacks& callbacks) {
  decoder_callbacks_ = &callbacks;
}

}  // namespace BackendRouting
}  // namespace HttpFilters
}  // namespace Extensions
}  // namespace Envoy
