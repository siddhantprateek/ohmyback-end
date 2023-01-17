package model

type Welcome7 struct {
	TrademarkApplication TrademarkApplication `json:"TrademarkApplication"`
}

type TrademarkApplication struct {
	RequestSearch      RequestSearch      `json:"RequestSearch"`
	RequestExamination RequestExamination `json:"RequestExamination"`
	TrademarkBag       TrademarkBag       `json:"TrademarkBag"`
	XmlnsCatmk         string             `json:"_xmlns:catmk"`
	XmlnsNs4           string             `json:"_xmlns:ns4"`
	XmlnsNs3           string             `json:"_xmlns:ns3"`
	XmlnsCOM           string             `json:"_xmlns:com"`
	XmlnsCacom         string             `json:"_xmlns:cacom"`
	XmlnsXs            string             `json:"_xmlns:xs"`
	XmlnsTmk           string             `json:"_xmlns:tmk"`
	XmlnsXsi           string             `json:"_xmlns:xsi"`
	COMSt96Version     string             `json:"_com:st96Version"`
	COMIPOVersion      string             `json:"_com:ipoVersion"`
	Prefix             Prefix             `json:"__prefix"`
}

type RequestExamination struct {
	RequestExaminationCategory RequestExaminationCategory `json:"RequestExaminationCategory"`
	Prefix                     Prefix                     `json:"__prefix"`
}

type RequestExaminationCategory struct {
	Prefix Prefix `json:"__prefix"`
	Text   string `json:"__text"`
}

type RequestSearch struct {
	RequestSearchCategory RequestExaminationCategory `json:"RequestSearchCategory"`
	Prefix                Prefix                     `json:"__prefix"`
}

type TrademarkBag struct {
	Trademark Trademark `json:"Trademark"`
	Prefix    Prefix    `json:"__prefix"`
}

type Trademark struct {
	RegistrationOfficeCode        RequestExaminationCategory   `json:"RegistrationOfficeCode"`
	ReceivingOfficeCode           RequestExaminationCategory   `json:"ReceivingOfficeCode"`
	ReceivingOfficeDate           RequestExaminationCategory   `json:"ReceivingOfficeDate"`
	ApplicationNumber             TrademarkApplicationNumber   `json:"ApplicationNumber"`
	RegistrationNumber            RequestExaminationCategory   `json:"RegistrationNumber"`
	ApplicationDate               RequestExaminationCategory   `json:"ApplicationDate"`
	RegistrationDate              RequestExaminationCategory   `json:"RegistrationDate"`
	FilingPlace                   RequestExaminationCategory   `json:"FilingPlace"`
	ApplicationLanguageCode       RequestExaminationCategory   `json:"ApplicationLanguageCode"`
	ExpiryDate                    RequestExaminationCategory   `json:"ExpiryDate"`
	MarkCurrentStatusCode         RequestExaminationCategory   `json:"MarkCurrentStatusCode"`
	MarkCurrentStatusDate         RequestExaminationCategory   `json:"MarkCurrentStatusDate"`
	MarkCategory                  RequestExaminationCategory   `json:"MarkCategory"`
	MarkRepresentation            MarkRepresentation           `json:"MarkRepresentation"`
	NonUseCancelledIndicator      RequestExaminationCategory   `json:"NonUseCancelledIndicator"`
	TradeDistinctivenessIndicator RequestExaminationCategory   `json:"TradeDistinctivenessIndicator"`
	UseRight                      []UseRight                   `json:"UseRight"`
	CommentText                   RequestExaminationCategory   `json:"CommentText"`
	OppositionPeriodStartDate     RequestExaminationCategory   `json:"OppositionPeriodStartDate"`
	OppositionPeriodEndDate       RequestExaminationCategory   `json:"OppositionPeriodEndDate"`
	GoodsServicesBag              GoodsServicesBag             `json:"GoodsServicesBag"`
	PriorityBag                   PriorityBag                  `json:"PriorityBag"`
	PublicationBag                PublicationBag               `json:"PublicationBag"`
	ApplicantBag                  ApplicantBag                 `json:"ApplicantBag"`
	NationalRepresentative        NationalRepresentative       `json:"NationalRepresentative"`
	MarkEventBag                  MarkEventBag                 `json:"MarkEventBag"`
	MarkFeatureDescription        RequestExaminationCategory   `json:"MarkFeatureDescription"`
	NationalTrademarkInformation  NationalTrademarkInformation `json:"NationalTrademarkInformation"`
	COMOperationCategory          string                       `json:"_com:operationCategory"`
	Prefix                        Prefix                       `json:"__prefix"`
}

type ApplicantBag struct {
	Applicant Applicant `json:"Applicant"`
	Prefix    Prefix    `json:"__prefix"`
}

type Applicant struct {
	LegalEntityName         RequestExaminationCategory `json:"LegalEntityName"`
	Contact                 ApplicantContact           `json:"Contact"`
	NationalLegalEntityCode RequestExaminationCategory `json:"NationalLegalEntityCode"`
	Prefix                  Prefix                     `json:"__prefix"`
}

type ApplicantContact struct {
	Name             Name                   `json:"Name"`
	PostalAddressBag PurplePostalAddressBag `json:"PostalAddressBag"`
	COMLanguageCode  COMLanguageCode        `json:"_com:languageCode"`
	Prefix           Prefix                 `json:"__prefix"`
}

type Name struct {
	EntityName UseRightText `json:"EntityName"`
	Prefix     Prefix       `json:"__prefix"`
}

type UseRightText struct {
	COMLanguageCode   *COMLanguageCode `json:"_com:languageCode,omitempty"`
	Prefix            Prefix           `json:"__prefix"`
	Text              string           `json:"__text"`
	COMSequenceNumber *string          `json:"_com:sequenceNumber,omitempty"`
}

type PurplePostalAddressBag struct {
	PostalAddress PurplePostalAddress `json:"PostalAddress"`
	Prefix        Prefix              `json:"__prefix"`
}

type PurplePostalAddress struct {
	PostalStructuredAddress PurplePostalStructuredAddress `json:"PostalStructuredAddress"`
	Prefix                  Prefix                        `json:"__prefix"`
}

type PurplePostalStructuredAddress struct {
	AddressLineText []UseRightText             `json:"AddressLineText"`
	CountryCode     RequestExaminationCategory `json:"CountryCode"`
	COMLanguageCode COMLanguageCode            `json:"_com:languageCode"`
	Prefix          Prefix                     `json:"__prefix"`
}

type TrademarkApplicationNumber struct {
	IPOfficeCode          RequestExaminationCategory `json:"IPOfficeCode"`
	ST13ApplicationNumber RequestExaminationCategory `json:"ST13ApplicationNumber"`
	Prefix                Prefix                     `json:"__prefix"`
}

type GoodsServicesBag struct {
	GoodsServices GoodsServices `json:"GoodsServices"`
	Prefix        Prefix        `json:"__prefix"`
}

type GoodsServices struct {
	GoodsServicesClassificationBag GoodsServicesClassificationBag   `json:"GoodsServicesClassificationBag"`
	NationalGoodsServices          NationalGoodsServices            `json:"NationalGoodsServices"`
	NationalFilingBasis            NationalFilingBasis              `json:"NationalFilingBasis"`
	ClassificationKindCode         RequestExaminationCategory       `json:"ClassificationKindCode"`
	ClassDescriptionBag            GoodsServicesClassDescriptionBag `json:"ClassDescriptionBag"`
	Prefix                         Prefix                           `json:"__prefix"`
}

type GoodsServicesClassDescriptionBag struct {
	ClassDescription []ClassDescription `json:"ClassDescription"`
	Prefix           Prefix             `json:"__prefix"`
}

type ClassDescription struct {
	GoodsServicesDescriptionText UseRightText `json:"GoodsServicesDescriptionText"`
	Prefix                       Prefix       `json:"__prefix"`
}

type GoodsServicesClassificationBag struct {
	GoodsServicesClassification []GoodsServicesClassification `json:"GoodsServicesClassification"`
	Prefix                      Prefix                        `json:"__prefix"`
}

type GoodsServicesClassification struct {
	ClassificationKindCode RequestExaminationCategory `json:"ClassificationKindCode"`
	CommentText            RequestExaminationCategory `json:"CommentText"`
	ClassNumber            RequestExaminationCategory `json:"ClassNumber"`
	ClassTitleText         RequestExaminationCategory `json:"ClassTitleText"`
	Prefix                 Prefix                     `json:"__prefix"`
}

type NationalFilingBasis struct {
	CurrentBasis CurrentBasis `json:"CurrentBasis"`
	Prefix       Prefix       `json:"__prefix"`
}

type CurrentBasis struct {
	BasisForeignApplicationIndicator  RequestExaminationCategory `json:"BasisForeignApplicationIndicator"`
	BasisForeignRegistrationIndicator RequestExaminationCategory `json:"BasisForeignRegistrationIndicator"`
	BasisUseIndicator                 RequestExaminationCategory `json:"BasisUseIndicator"`
	BasisIntentToUseIndicator         RequestExaminationCategory `json:"BasisIntentToUseIndicator"`
	NoBasisIndicator                  RequestExaminationCategory `json:"NoBasisIndicator"`
	Prefix                            Prefix                     `json:"__prefix"`
}

type NationalGoodsServices struct {
	NationalClassTotalQuantity RequestExaminationCategory `json:"NationalClassTotalQuantity"`
	Prefix                     Prefix                     `json:"__prefix"`
}

type MarkEventBag struct {
	MarkEvent []MarkEvent `json:"MarkEvent"`
	Prefix    Prefix      `json:"__prefix"`
}

type MarkEvent struct {
	MarkEventCategory     RequestExaminationCategory  `json:"MarkEventCategory"`
	NationalMarkEvent     NationalMarkEvent           `json:"NationalMarkEvent"`
	MarkEventDate         RequestExaminationCategory  `json:"MarkEventDate"`
	Prefix                Prefix                      `json:"__prefix"`
	MarkEventResponseDate *RequestExaminationCategory `json:"MarkEventResponseDate,omitempty"`
}

type NationalMarkEvent struct {
	MarkEventCode                            RequestExaminationCategory               `json:"MarkEventCode"`
	MarkEventDescriptionText                 RequestExaminationCategory               `json:"MarkEventDescriptionText"`
	MarkEventOtherLanguageDescriptionTextBag MarkEventOtherLanguageDescriptionTextBag `json:"MarkEventOtherLanguageDescriptionTextBag"`
	XsiType                                  XsiType                                  `json:"_xsi:type"`
	Prefix                                   Prefix                                   `json:"__prefix"`
	MarkEventAdditionalText                  *RequestExaminationCategory              `json:"MarkEventAdditionalText,omitempty"`
}

type MarkEventOtherLanguageDescriptionTextBag struct {
	MarkEventDescriptionText UseRightText `json:"MarkEventDescriptionText"`
	Prefix                   Prefix       `json:"__prefix"`
}

type MarkRepresentation struct {
	MarkFeatureCategory RequestExaminationCategory `json:"MarkFeatureCategory"`
	MarkReproduction    MarkReproduction           `json:"MarkReproduction"`
	Prefix              Prefix                     `json:"__prefix"`
}

type MarkReproduction struct {
	MarkImageBag MarkImageBag `json:"MarkImageBag"`
	Prefix       Prefix       `json:"__prefix"`
}

type MarkImageBag struct {
	MarkImage MarkImage `json:"MarkImage"`
	Prefix    Prefix    `json:"__prefix"`
}

type MarkImage struct {
	FileName                RequestExaminationCategory `json:"FileName"`
	ImageFormatCategory     RequestExaminationCategory `json:"ImageFormatCategory"`
	MarkImageClassification MarkImageClassification    `json:"MarkImageClassification"`
	Prefix                  Prefix                     `json:"__prefix"`
}

type MarkImageClassification struct {
	FigurativeElementClassificationBag FigurativeElementClassificationBag `json:"FigurativeElementClassificationBag"`
	Prefix                             Prefix                             `json:"__prefix"`
}

type FigurativeElementClassificationBag struct {
	ViennaClassificationBag ViennaClassificationBag `json:"ViennaClassificationBag"`
	Prefix                  Prefix                  `json:"__prefix"`
}

type ViennaClassificationBag struct {
	ViennaClassification []ViennaClassification `json:"ViennaClassification"`
	Prefix               Prefix                 `json:"__prefix"`
}

type ViennaClassification struct {
	ViennaCategory       RequestExaminationCategory `json:"ViennaCategory"`
	ViennaDivision       RequestExaminationCategory `json:"ViennaDivision"`
	ViennaSection        RequestExaminationCategory `json:"ViennaSection"`
	ViennaDescriptionBag ViennaDescriptionBag       `json:"ViennaDescriptionBag"`
	XsiType              string                     `json:"_xsi:type"`
	Prefix               Prefix                     `json:"__prefix"`
}

type ViennaDescriptionBag struct {
	ViennaDescription []UseRightText `json:"ViennaDescription"`
	Prefix            Prefix         `json:"__prefix"`
}

type NationalRepresentative struct {
	CommentText RequestExaminationCategory    `json:"CommentText"`
	Contact     NationalRepresentativeContact `json:"Contact"`
	Prefix      Prefix                        `json:"__prefix"`
}

type NationalRepresentativeContact struct {
	Name             Name                   `json:"Name"`
	PostalAddressBag FluffyPostalAddressBag `json:"PostalAddressBag"`
	COMLanguageCode  COMLanguageCode        `json:"_com:languageCode"`
	Prefix           Prefix                 `json:"__prefix"`
}

type FluffyPostalAddressBag struct {
	PostalAddress FluffyPostalAddress `json:"PostalAddress"`
	Prefix        Prefix              `json:"__prefix"`
}

type FluffyPostalAddress struct {
	PostalStructuredAddress FluffyPostalStructuredAddress `json:"PostalStructuredAddress"`
	Prefix                  Prefix                        `json:"__prefix"`
}

type FluffyPostalStructuredAddress struct {
	AddressLineText      []UseRightText             `json:"AddressLineText"`
	GeographicRegionName GeographicRegionName       `json:"GeographicRegionName"`
	CountryCode          RequestExaminationCategory `json:"CountryCode"`
	PostalCode           RequestExaminationCategory `json:"PostalCode"`
	COMLanguageCode      COMLanguageCode            `json:"_com:languageCode"`
	Prefix               Prefix                     `json:"__prefix"`
}

type GeographicRegionName struct {
	COMGeographicRegionCategory string `json:"_com:geographicRegionCategory"`
	Prefix                      Prefix `json:"__prefix"`
	Text                        string `json:"__text"`
}

type NationalTrademarkInformation struct {
	RegisterCategory                         RequestExaminationCategory `json:"RegisterCategory"`
	MarkCurrentStatusInternalDescriptionText RequestExaminationCategory `json:"MarkCurrentStatusInternalDescriptionText"`
	RenewalDate                              RequestExaminationCategory `json:"RenewalDate"`
	AllowedDate                              RequestExaminationCategory `json:"AllowedDate"`
	TrademarkClass                           TrademarkClass             `json:"TrademarkClass"`
	Legislation                              Legislation                `json:"Legislation"`
	InterestedPartyBag                       InterestedPartyBag         `json:"InterestedPartyBag"`
	ClaimBag                                 ClaimBag                   `json:"ClaimBag"`
	MarkFeatureCategoryBag                   MarkFeatureCategoryBag     `json:"MarkFeatureCategoryBag"`
	XsiType                                  string                     `json:"_xsi:type"`
	Prefix                                   Prefix                     `json:"__prefix"`
}

type ClaimBag struct {
	Claim  []Claim `json:"Claim"`
	Prefix Prefix  `json:"__prefix"`
}

type Claim struct {
	ClaimCategoryType           RequestExaminationCategory  `json:"ClaimCategoryType"`
	ClaimTypeDescription        UseRightText                `json:"ClaimTypeDescription"`
	ClaimNumber                 RequestExaminationCategory  `json:"ClaimNumber"`
	ClaimDate                   *ClaimDate                  `json:"ClaimDate,omitempty"`
	ClaimCountryCode            *RequestExaminationCategory `json:"ClaimCountryCode,omitempty"`
	ClaimForeignRegistrationNbr *RequestExaminationCategory `json:"ClaimForeignRegistrationNbr,omitempty"`
	ClaimText                   RequestExaminationCategory  `json:"ClaimText"`
	GoodsServicesReferenceBag   GoodsServicesReferenceBag   `json:"GoodsServicesReferenceBag"`
	Prefix                      Prefix                      `json:"__prefix"`
}

type ClaimDate struct {
	StructuredClaimDate RequestExaminationCategory `json:"StructuredClaimDate"`
	Prefix              Prefix                     `json:"__prefix"`
}

type GoodsServicesReferenceBag struct {
	GoodsServicesReference *GoodsServicesReferenceUnion `json:"GoodsServicesReference"`
	Prefix                 Prefix                       `json:"__prefix"`
}

type GoodsServicesReferenceElement struct {
	GoodsServicesReferenceIdentifier RequestExaminationCategory `json:"GoodsServicesReferenceIdentifier"`
	Prefix                           Prefix                     `json:"__prefix"`
}

type InterestedPartyBag struct {
	InterestedParty []InterestedParty `json:"InterestedParty"`
	Prefix          Prefix            `json:"__prefix"`
}

type InterestedParty struct {
	InterestedPartyCategory RequestExaminationCategory `json:"InterestedPartyCategory"`
	Contact                 ApplicantContact           `json:"Contact"`
	Prefix                  Prefix                     `json:"__prefix"`
}

type Legislation struct {
	LegislationCode           RequestExaminationCategory `json:"LegislationCode"`
	LegislationDescriptionBag LegislationDescriptionBag  `json:"LegislationDescriptionBag"`
	Prefix                    Prefix                     `json:"__prefix"`
}

type LegislationDescriptionBag struct {
	LegislationDescription []UseRightText `json:"LegislationDescription"`
	Prefix                 Prefix         `json:"__prefix"`
}

type MarkFeatureCategoryBag struct {
	MarkFeatureCategory RequestExaminationCategory `json:"MarkFeatureCategory"`
	Prefix              Prefix                     `json:"__prefix"`
}

type TrademarkClass struct {
	TrademarkClassCode           RequestExaminationCategory   `json:"TrademarkClassCode"`
	TrademarkClassDescriptionBag TrademarkClassDescriptionBag `json:"TrademarkClassDescriptionBag"`
	Prefix                       Prefix                       `json:"__prefix"`
}

type TrademarkClassDescriptionBag struct {
	TrademarkClassDescription []UseRightText `json:"TrademarkClassDescription"`
	Prefix                    Prefix         `json:"__prefix"`
}

type PriorityBag struct {
	Priority []Priority `json:"Priority"`
	Prefix   Prefix     `json:"__prefix"`
}

type Priority struct {
	PriorityCountryCode           RequestExaminationCategory   `json:"PriorityCountryCode"`
	ApplicationNumber             PriorityApplicationNumber    `json:"ApplicationNumber"`
	PriorityApplicationFilingDate RequestExaminationCategory   `json:"PriorityApplicationFilingDate"`
	PriorityPartialGoodsServices  PriorityPartialGoodsServices `json:"PriorityPartialGoodsServices"`
	Prefix                        Prefix                       `json:"__prefix"`
}

type PriorityApplicationNumber struct {
	ApplicationNumberText RequestExaminationCategory `json:"ApplicationNumberText"`
	Prefix                Prefix                     `json:"__prefix"`
}

type PriorityPartialGoodsServices struct {
	ClassDescriptionBag PriorityPartialGoodsServicesClassDescriptionBag `json:"ClassDescriptionBag"`
	Prefix              Prefix                                          `json:"__prefix"`
}

type PriorityPartialGoodsServicesClassDescriptionBag struct {
	ClassDescription ClassDescription `json:"ClassDescription"`
	Prefix           Prefix           `json:"__prefix"`
}

type PublicationBag struct {
	Publication Publication `json:"Publication"`
	Prefix      Prefix      `json:"__prefix"`
}

type Publication struct {
	PublicationIdentifier RequestExaminationCategory `json:"PublicationIdentifier"`
	NationalPublication   NationalPublication        `json:"NationalPublication"`
	Prefix                Prefix                     `json:"__prefix"`
}

type NationalPublication struct {
	PublicationStatusCategory RequestExaminationCategory `json:"PublicationStatusCategory"`
	PublicationActionDate     RequestExaminationCategory `json:"PublicationActionDate"`
	Prefix                    Prefix                     `json:"__prefix"`
}

type UseRight struct {
	UseRightIndicator RequestExaminationCategory `json:"UseRightIndicator"`
	UseRightText      UseRightText               `json:"UseRightText"`
	Prefix            Prefix                     `json:"__prefix"`
}

type Prefix string

const (
	COM   Prefix = "com"
	Cacom Prefix = "cacom"
	Catmk Prefix = "catmk"
	Tmk   Prefix = "tmk"
)

type COMLanguageCode string

const (
	En COMLanguageCode = "en"
	Fr COMLanguageCode = "fr"
)

type XsiType string

const (
	CatmkNationalMarkEventType XsiType = "catmk:NationalMarkEventType"
)

type GoodsServicesReferenceUnion struct {
	GoodsServicesReferenceElement      *GoodsServicesReferenceElement
	GoodsServicesReferenceElementArray []GoodsServicesReferenceElement
}
