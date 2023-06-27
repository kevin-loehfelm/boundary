// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package warning

import pb "github.com/hashicorp/boundary/internal/gen/controller/api"

type apiWarning uint

// The set of warnings that boundary ever returns as a result of API requests.
// Besides zzzKeepThisLastSentinel, the warnings should keep the numbers they
// are initially released with because the enumerated number is used to uniquely
// identify them and potentially provide additional information in documentation.
const (
	Unknown                                   apiWarning = 0
	FieldDeprecatedTargetWorkerFilters        apiWarning = 1
	ActionDeprecatedAddCredentialLibraries    apiWarning = 2
	ActionDeprecatedSetCredentialLibraries    apiWarning = 3
	ActionDeprecatedRemoveCredentialLibraries apiWarning = 4
	OidcAuthMethodInactiveCannotBeUsed        apiWarning = 5

	// This is a sentinel value that captures the largest apiWarning id currently
	// known.  Add all warnings above this line.
	zzzKeepThisLastSentinel
)

func (a apiWarning) toProto() *pb.Warning {
	nw := &pb.Warning{
		Code: int32(a),
	}
	switch a {
	case FieldDeprecatedTargetWorkerFilters:
		nw.Warning = &pb.Warning_RequestField{RequestField: &pb.FieldWarning{
			Name:    "worker_filter",
			Warning: "This field is deprecated. Please use ingress_worker_filter and/or egress_worker_filter",
		}}
	case ActionDeprecatedAddCredentialLibraries:
		nw.Warning = &pb.Warning_Action{Action: &pb.ActionWarning{
			Name:    "add-credential-libraries",
			Warning: "This action is deprecated. Use ':add-credential-sources' instead.",
		}}
	case ActionDeprecatedSetCredentialLibraries:
		nw.Warning = &pb.Warning_Action{Action: &pb.ActionWarning{
			Name:    "set-credential-libraries",
			Warning: "This action is deprecated. Use ':set-credential-sources' instead.",
		}}
	case ActionDeprecatedRemoveCredentialLibraries:
		nw.Warning = &pb.Warning_Action{Action: &pb.ActionWarning{
			Name:    "remove-credential-libraries",
			Warning: "This action is deprecated. Use ':remove-credential-sources' instead.",
		}}
	case OidcAuthMethodInactiveCannotBeUsed:
		nw.Warning = &pb.Warning_Behavior{Behavior: &pb.BehaviorWarning{
			Warning: "OIDC Auth Methods cannot be authenticated until they have been made active.",
		}}
	default:
		// don't add any unknown warning to the warner
		return nil
	}
	return nw
}
