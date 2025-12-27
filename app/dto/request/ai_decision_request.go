package request

import "backend-dinakom/app/constants"

type CreateDecisionRequest struct {
	NeedSearch  bool                  `json:"needSearch" binding:"required"`
	Queries     []string              `json:"queries" binding:"required_if=NeedSearch true,dive,required"`
	RiskLevel   constants.AIRiskLevel `json:"riskLevel" binding:"required,oneof=LOW MEDIUM HIGH"`
	RequestType string                `json:"requestType" binding:"required,oneof=GOOGLE_SEARCH EXTRACT ANSWER"`
}

type UpdateDecisionRequest struct {
	NeedSearch  *bool                  `json:"needSearch" binding:"required"`
	Queries     []string              `json:"queries" binding:"required_if=NeedSearch true,dive,required"`
	RiskLevel   *constants.AIRiskLevel `json:"riskLevel" binding:"required,oneof=LOW MEDIUM HIGH"`
	RequestType *string                `json:"requestType" binding:"required,oneof=GOOGLE_SEARCH EXTRACT ANSWER"`
}
