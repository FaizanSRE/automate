// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/cfgmgmt/cfgmgmt.proto

package cfgmgmt

import (
	request "github.com/chef/automate/api/external/cfgmgmt/request"
	query "github.com/chef/automate/api/external/common/query"
	policy "github.com/chef/automate/api/external/iam/v2/policy"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodes", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/nodes", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Nodes); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "start":
					return m.Start
				case "end":
					return m.End
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetRuns", "infra:nodes:{node_id}", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/nodes/{node_id}/runs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Runs); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				case "start":
					return m.Start
				case "end":
					return m.End
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodesCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/node_counts", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodesCounts); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "start":
					return m.Start
				case "end":
					return m.End
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetRunsCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/run_counts", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RunsCounts); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "start":
					return m.Start
				case "end":
					return m.End
				case "node_id":
					return m.NodeId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetCheckInCountsTimeSeries", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/checkin_counts_timeseries", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetMissingNodeDurationCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/stats/missing_node_duration_counts", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeRun", "infra:nodes:{node_id}", "infra:nodes:get", "GET", "/api/v0/cfgmgmt/nodes/{node_id}/runs/{run_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeRun); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				case "run_id":
					return m.RunId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetSuggestions", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/suggestions", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*query.Suggestion); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "type":
					return m.Type
				case "text":
					return m.Text
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetOrganizations", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/organizations", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetSourceFqdns", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/source_fqdns", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetAttributes", "infra:nodes:{node_id}", "infra:nodes:get", "GET", "/api/v0/cfgmgmt/nodes/{node_id}/attribute", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Node); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/api/v0/cfgmgmt/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetPolicyCookbooks", "infra:nodes:{revision_id}", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/policy_revision/{revision_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.PolicyRevision); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "revision_id":
					return m.RevisionId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetErrors", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/errors", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeMetadataCounts", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/node_metadata_counts", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeMetadataCounts); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "start":
					return m.Start
				case "end":
					return m.End
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetNodeRunsDailyStatusTimeSeries", "infra:nodes", "infra:nodes:list", "GET", "/api/v0/cfgmgmt/node_runs_daily_status_time_series", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.NodeRunsDailyStatusTimeSeries); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "node_id":
					return m.NodeId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/CreateRollout", "ingest:unifiedEvents", "ingest:unifiedEvents:create", "POST", "/api/beta/cfgmgmt/rollouts/create", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateRollout); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "policy_name":
					return m.PolicyName
				case "policy_node_group":
					return m.PolicyNodeGroup
				case "policy_revision_id":
					return m.PolicyRevisionId
				case "policy_domain_url":
					return m.PolicyDomainUrl
				case "policy_scm_url":
					return m.PolicyScmUrl
				case "policy_scm_web_url":
					return m.PolicyScmWebUrl
				case "policy_scm_commit":
					return m.PolicyScmCommit
				case "description":
					return m.Description
				case "ci_job_url":
					return m.CiJobUrl
				case "ci_job_id":
					return m.CiJobId
				case "scm_author_name":
					return m.ScmAuthorName
				case "scm_author_email":
					return m.ScmAuthorEmail
				case "policy_domain_username":
					return m.PolicyDomainUsername
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/CreateRolloutTest", "ingest:unifiedEvents", "ingest:unifiedEvents:create", "POST", "/api/beta/cfgmgmt/rollouts/test_create", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetRollouts", "infra:nodes", "infra:nodes:list", "GET", "/api/beta/cfgmgmt/rollouts/list", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetRolloutById", "infra:nodes", "infra:nodes:list", "GET", "/api/beta/cfgmgmt/rollouts/rollout/{rollout_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RolloutById); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "rollout_id":
					return m.RolloutId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/GetRolloutForChefRun", "infra:nodes", "infra:nodes:list", "GET", "/api/beta/cfgmgmt/rollouts/find", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RolloutForChefRun); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "policy_name":
					return m.PolicyName
				case "policy_group":
					return m.PolicyGroup
				case "policy_revision_id":
					return m.PolicyRevisionId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.cfgmgmt.ConfigMgmt/ListNodeSegmentsWithRolloutProgress", "infra:nodes", "infra:nodes:list", "GET", "/api/beta/cfgmgmt/rollouts/progress_by_node_segment", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
