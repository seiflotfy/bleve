//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package stop_words_filter

import (
	"github.com/couchbaselabs/bleve/analysis"
)

type StopWordsFilter struct {
	stopWords analysis.WordMap
}

func NewStopWordsFilter(stopWords analysis.WordMap) *StopWordsFilter {
	return &StopWordsFilter{
		stopWords: stopWords,
	}
}

func (f *StopWordsFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	rv := make(analysis.TokenStream, 0)

	for _, token := range input {
		word := string(token.Term)
		_, isStopWord := f.stopWords[word]
		if !isStopWord {
			rv = append(rv, token)
		}
	}

	return rv
}
