//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package simple

import (
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/xpipeline"
)

type SimpleExecutablePipelineBuilder struct {
	site        catalog.Site
	defaultPool string
}

func NewSimpleExecutablePipelineBuilder(site catalog.Site, defaultPool string) *SimpleExecutablePipelineBuilder {
	return &SimpleExecutablePipelineBuilder{
		site:        site,
		defaultPool: defaultPool,
	}
}

func (this *SimpleExecutablePipelineBuilder) Build(p *plan.Plan) (*xpipeline.ExecutablePipeline, error) {
	rv := &xpipeline.ExecutablePipeline{}

	var lastOperator xpipeline.Operator = nil
	currentElement := p.Root

	for currentElement != nil {
		var currentOperator xpipeline.Operator = nil
		switch currentElement := currentElement.(type) {
		case *plan.Scan:
			pool, err := this.site.Pool(currentElement.Pool)
			if err != nil {
				return nil, err
			}
			bucket, err := pool.Bucket(currentElement.Bucket)
			if err != nil {
				return nil, err
			}
			scanner, err := bucket.Scanner(currentElement.Scanner)
			if err != nil {
				return nil, err
			}
			currentOperator = xpipeline.NewScan(bucket, scanner)
		case *plan.ExpressionEvaluator:
			currentOperator = xpipeline.NewExpressionEvaluatorSource()
		case *plan.Fetch:
			pool, err := this.site.Pool(currentElement.Pool)
			if err != nil {
				return nil, err
			}
			bucket, err := pool.Bucket(currentElement.Bucket)
			if err != nil {
				return nil, err
			}
			currentOperator = xpipeline.NewFetch(bucket, currentElement.Projection, currentElement.As)
		case *plan.Filter:
			currentOperator = xpipeline.NewFilter(currentElement.Expr)
		case *plan.Order:
			currentOperator = xpipeline.NewOrder(currentElement.Sort, currentElement.ExplicitAliases)
		case *plan.Limit:
			currentOperator = xpipeline.NewLimit(currentElement.Val)
		case *plan.Offset:
			currentOperator = xpipeline.NewOffset(currentElement.Val)
		case *plan.Projector:
			currentOperator = xpipeline.NewProject(currentElement.Result, currentElement.ProjectEmpty)
		case *plan.DocumentJoin:
			currentOperator = xpipeline.NewDocumentJoin(currentElement.Over, currentElement.As)
		case *plan.EliminateDuplicates:
			currentOperator = xpipeline.NewEliminateDuplicates()
		case *plan.Grouper:
			currentOperator = xpipeline.NewGrouper(currentElement.Group, currentElement.Aggregates)
		}

		//link root of xpipeline
		if rv.Root == nil {
			rv.Root = currentOperator
		}

		// link previous operator to this one
		if lastOperator != nil {
			lastOperator.SetSource(currentOperator)
		}

		// advance to next
		lastOperator = currentOperator

		sources := currentElement.Sources()
		if len(sources) > 1 {
			// FIXME future operators like JOIN will have more than one source
			clog.Fatalf("multiple sources not yet supported")
		} else if len(sources) == 1 {
			currentElement = sources[0]
		} else {
			currentElement = nil
		}
	}

	return rv, nil
}
