package finvoice

import (
	"bytes"
	"encoding/xml"
	"time"

	"github.com/juelko/invoice/pkg/message"
)

type Amount struct {
	MonetaryAmount           MonetaryAmount           `xml:",chardata"`
	AmountCurrencyIdentifier AmountCurrencyIdentifier `xml:"AmountCurrencyIdentifier,attr"`
}

// May be no more than 3 items long
type AmountCurrencyIdentifier string

type AnyPartyCommunicationDetailsType struct {
	AnyPartyPhoneNumberIdentifier  string `xml:"AnyPartyPhoneNumberIdentifier,omitempty"`
	AnyPartyEmailAddressIdentifier string `xml:"AnyPartyEmailAddressIdentifier,omitempty"`
}

type AnyPartyDetailsType struct {
	AnyPartyText                    Anypartytexttype035              `xml:"AnyPartyText"`
	AnyPartyIdentifier              PartyLegalRegIdType              `xml:"AnyPartyIdentifier,omitempty"`
	AnyPartyOrganisationName        []GenericStringType235           `xml:"AnyPartyOrganisationName"`
	AnyPartyOrganisationDepartment  []string                         `xml:"AnyPartyOrganisationDepartment,omitempty"`
	AnyPartyOrganisationTaxCode     string                           `xml:"AnyPartyOrganisationTaxCode,omitempty"`
	AnyPartyCode                    PartyIdentifierType              `xml:"AnyPartyCode,omitempty"`
	AnyPartyContactPersonName       string                           `xml:"AnyPartyContactPersonName,omitempty"`
	AnyPartyContactPersonFunction   []string                         `xml:"AnyPartyContactPersonFunction,omitempty"`
	AnyPartyContactPersonDepartment []string                         `xml:"AnyPartyContactPersonDepartment,omitempty"`
	AnyPartyCommunicationDetails    AnyPartyCommunicationDetailsType `xml:"AnyPartyCommunicationDetails,omitempty"`
	AnyPartyPostalAddressDetails    AnyPartyPostalAddressDetails     `xml:"AnyPartyPostalAddressDetails,omitempty"`
	AnyPartyOrganisationUnitNumber  string                           `xml:"AnyPartyOrganisationUnitNumber,omitempty"`
	AnyPartySiteCode                string                           `xml:"AnyPartySiteCode,omitempty"`
}

type AnyPartyPostalAddressDetails struct {
	AnyPartyStreetName              []GenericStringType235 `xml:"AnyPartyStreetName"`
	AnyPartyTownName                GenericStringType235   `xml:"AnyPartyTownName"`
	AnyPartyPostCodeIdentifier      GenericStringType235   `xml:"AnyPartyPostCodeIdentifier"`
	AnyPartyCountrySubdivision      GenericStringType235   `xml:"AnyPartyCountrySubdivision,omitempty"`
	CountryCode                     CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                     string                 `xml:"CountryName,omitempty"`
	AnyPartyPostOfficeBoxIdentifier string                 `xml:"AnyPartyPostOfficeBoxIdentifier,omitempty"`
}

type AnyPartyTextType struct {
	Value        string `xml:",chardata"`
	AnyPartyCode string `xml:"AnyPartyCode,attr"`
}

type Anypartytexttype035 struct {
	Value        string `xml:",chardata"`
	AnyPartyCode string `xml:"AnyPartyCode,attr"`
}

type ArticleGroupIdentifierType struct {
	Value         string               `xml:",chardata"`
	SchemeID      string               `xml:"SchemeID,attr,omitempty"`
	SchemeVersion GenericStringType135 `xml:"SchemeVersion,attr,omitempty"`
}

type AttachmentMessageDetailsType struct {
	AttachmentMessageIdentifier AttachmentsIdentifierType `xml:"AttachmentMessageIdentifier"`
}

// Must match the pattern .{2,48}::attachments
type AttachmentsIdentifierType string

type BuyerCommunicationDetailsType struct {
	BuyerPhoneNumberIdentifier  string `xml:"BuyerPhoneNumberIdentifier,omitempty"`
	BuyerEmailaddressIdentifier string `xml:"BuyerEmailaddressIdentifier,omitempty"`
}

type BuyerPartyDetailsType struct {
	BuyerPartyIdentifier         PartyLegalRegIdType           `xml:"BuyerPartyIdentifier,omitempty"`
	BuyerOrganisationName        []GenericStringType270        `xml:"BuyerOrganisationName"`
	BuyerOrganisationTradingName GenericStringType270          `xml:"BuyerOrganisationTradingName,omitempty"`
	BuyerOrganisationDepartment  []string                      `xml:"BuyerOrganisationDepartment,omitempty"`
	BuyerOrganisationTaxCode     string                        `xml:"BuyerOrganisationTaxCode,omitempty"`
	BuyerCode                    PartyIdentifierType           `xml:"BuyerCode,omitempty"`
	BuyerPostalAddressDetails    BuyerPostalAddressDetailsType `xml:"BuyerPostalAddressDetails,omitempty"`
}

type BuyerPostalAddressDetailsType struct {
	BuyerStreetName              []GenericStringType235 `xml:"BuyerStreetName"`
	BuyerTownName                GenericStringType235   `xml:"BuyerTownName"`
	BuyerPostCodeIdentifier      GenericStringType235   `xml:"BuyerPostCodeIdentifier"`
	BuyerCountrySubdivision      GenericStringType235   `xml:"BuyerCountrySubdivision,omitempty"`
	CountryCode                  CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                  string                 `xml:"CountryName,omitempty"`
	BuyerPostOfficeBoxIdentifier string                 `xml:"BuyerPostOfficeBoxIdentifier,omitempty"`
}

// Must match the pattern [A-Z0-9]*
type CapAZ09 string

type CashDiscountVatDetails struct {
	CashDiscountVatPercent Percentage `xml:"CashDiscountVatPercent"`
	CashDiscountVatAmount  Amount     `xml:"CashDiscountVatAmount"`
}

// May be one of SHA, OUR, BEN, SLEV
type ChargeOption string

// May be one of SPY
type CodeListAgencyIdentifier string

// May be no more than 2 items long
type CountryCodeType string

type CustomsInfoType struct {
	CNCode                     GenericStringType18  `xml:"CNCode,omitempty"`
	CNName                     GenericStringType135 `xml:"CNName,omitempty"`
	CNOriginCountrySubdivision GenericStringType235 `xml:"CNOriginCountrySubdivision,omitempty"`
	CNOriginCountryCode        CountryCodeType      `xml:"CNOriginCountryCode,omitempty"`
	CNOriginCountryName        GenericStringType135 `xml:"CNOriginCountryName,omitempty"`
}

type Date struct {
	DateType DateType `xml:",chardata"`
	Format   Format   `xml:"Format,attr"`
}

// Must match the pattern [0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]
type DateType int

type DefinitionDetails struct {
	DefinitionHeaderText DefinitionHeaderText `xml:"DefinitionHeaderText"`
	DefinitionValue      QuantityType070      `xml:"DefinitionValue,omitempty"`
}

type DefinitionHeaderText struct {
	Value          string              `xml:",chardata"`
	DefinitionCode GenericTokenType120 `xml:"DefinitionCode,attr,omitempty"`
}

type DeliveryCommunicationDetailsType struct {
	DeliveryPhoneNumberIdentifier  string `xml:"DeliveryPhoneNumberIdentifier,omitempty"`
	DeliveryEmailaddressIdentifier string `xml:"DeliveryEmailaddressIdentifier,omitempty"`
}

type DeliveryDetailsType struct {
	DeliveryDate                   Date                      `xml:"DeliveryDate,omitempty"`
	DeliveryPeriodDetails          DeliveryPeriodDetailsType `xml:"DeliveryPeriodDetails,omitempty"`
	ShipmentPartyDetails           ShipmentPartyDetails      `xml:"ShipmentPartyDetails,omitempty"`
	DeliveryMethodText             string                    `xml:"DeliveryMethodText,omitempty"`
	DeliveryTermsText              string                    `xml:"DeliveryTermsText,omitempty"`
	DeliveryTermsCode              GenericStringType14       `xml:"DeliveryTermsCode,omitempty"`
	TerminalAddressText            string                    `xml:"TerminalAddressText,omitempty"`
	WaybillIdentifier              string                    `xml:"WaybillIdentifier,omitempty"`
	WaybillTypeCode                string                    `xml:"WaybillTypeCode,omitempty"`
	ClearanceIdentifier            string                    `xml:"ClearanceIdentifier,omitempty"`
	DeliveryNoteIdentifier         string                    `xml:"DeliveryNoteIdentifier,omitempty"`
	DelivererIdentifier            string                    `xml:"DelivererIdentifier,omitempty"`
	DelivererName                  []string                  `xml:"DelivererName,omitempty"`
	DelivererCountrySubdivision    GenericStringType235      `xml:"DelivererCountrySubdivision,omitempty"`
	DelivererCountryCode           CountryCodeType           `xml:"DelivererCountryCode,omitempty"`
	DelivererCountryName           string                    `xml:"DelivererCountryName,omitempty"`
	ModeOfTransportIdentifier      string                    `xml:"ModeOfTransportIdentifier,omitempty"`
	CarrierName                    string                    `xml:"CarrierName,omitempty"`
	VesselName                     string                    `xml:"VesselName,omitempty"`
	LocationIdentifier             string                    `xml:"LocationIdentifier,omitempty"`
	TransportInformationDate       Date                      `xml:"TransportInformationDate,omitempty"`
	CountryOfOrigin                string                    `xml:"CountryOfOrigin,omitempty"`
	CountryOfDestinationName       string                    `xml:"CountryOfDestinationName,omitempty"`
	DestinationCountryCode         CountryCodeType           `xml:"DestinationCountryCode,omitempty"`
	PlaceOfDischarge               []string                  `xml:"PlaceOfDischarge,omitempty"`
	FinalDestinationName           []DestinationNameType     `xml:"FinalDestinationName,omitempty"`
	ManufacturerIdentifier         string                    `xml:"ManufacturerIdentifier,omitempty"`
	ManufacturerName               []string                  `xml:"ManufacturerName,omitempty"`
	ManufacturerCountrySubdivision GenericStringType235      `xml:"ManufacturerCountrySubdivision,omitempty"`
	ManufacturerCountryCode        CountryCodeType           `xml:"ManufacturerCountryCode,omitempty"`
	ManufacturerCountryName        string                    `xml:"ManufacturerCountryName,omitempty"`
	ManufacturerOrderIdentifier    string                    `xml:"ManufacturerOrderIdentifier,omitempty"`
	PackageDetails                 PackageDetails            `xml:"PackageDetails,omitempty"`
}

type DeliveryPartyDetailsType struct {
	DeliveryPartyIdentifier        PartyLegalRegIdType              `xml:"DeliveryPartyIdentifier,omitempty"`
	DeliveryOrganisationName       []GenericStringType235           `xml:"DeliveryOrganisationName"`
	DeliveryOrganisationDepartment []string                         `xml:"DeliveryOrganisationDepartment,omitempty"`
	DeliveryOrganisationTaxCode    string                           `xml:"DeliveryOrganisationTaxCode,omitempty"`
	DeliveryCode                   PartyIdentifierType              `xml:"DeliveryCode,omitempty"`
	DeliveryPostalAddressDetails   DeliveryPostalAddressDetailsType `xml:"DeliveryPostalAddressDetails"`
}

type DeliveryPeriodDetailsType struct {
	StartDate Date `xml:"StartDate"`
	EndDate   Date `xml:"EndDate"`
}

type DeliveryPostalAddressDetailsType struct {
	DeliveryStreetName              []GenericStringType235 `xml:"DeliveryStreetName"`
	DeliveryTownName                GenericStringType235   `xml:"DeliveryTownName"`
	DeliveryPostCodeIdentifier      GenericStringType235   `xml:"DeliveryPostCodeIdentifier"`
	DeliveryCountrySubdivision      GenericStringType235   `xml:"DeliveryCountrySubdivision,omitempty"`
	CountryCode                     CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                     string                 `xml:"CountryName,omitempty"`
	DeliveryPostofficeBoxIdentifier string                 `xml:"DeliveryPostofficeBoxIdentifier,omitempty"`
}

type DestinationNameType struct {
	Value    string     `xml:",chardata"`
	SchemeID Iso6523cid `xml:"SchemeID,attr,omitempty"`
}

type DirectDebitInfoType struct {
	MandateReference   GenericStringType135 `xml:"MandateReference,omitempty"`
	CreditorIdentifier GenericStringType135 `xml:"CreditorIdentifier,omitempty"`
	DebitedAccountID   EpiAccountIDType     `xml:"DebitedAccountID,omitempty"`
}

type DiscountDetailsType struct {
	FreeText        GenericStringType170 `xml:"FreeText,omitempty"`
	ReasonCode      string               `xml:"ReasonCode,omitempty"`
	Percent         Percentage           `xml:"Percent,omitempty"`
	Amount          Amount               `xml:"Amount,omitempty"`
	BaseAmount      Amount               `xml:"BaseAmount,omitempty"`
	VatCategoryCode string               `xml:"VatCategoryCode,omitempty"`
	VatRatePercent  Percentage           `xml:"VatRatePercent,omitempty"`
}

type EanCodeType struct {
	Value    string     `xml:",chardata"`
	SchemeID Iso6523cid `xml:"SchemeID,attr,omitempty"`
}

type FinvoiceAddress struct {
	Address  AddressID       `xml:",chardata"`
	SchemeID AddressSchemeID `xml:"SchemeID,attr,omitempty"`
}

type AddressID string

// Must be at least 1 items long
type AddressSchemeID string

type EpiAccountIDType struct {
	GenericNMtokenType134    GenericNMtokenType134    `xml:",chardata"`
	IdentificationSchemeName IdentificationSchemeName `xml:"IdentificationSchemeName,attr"`
}

type EpiAmount struct {
	EpiMonetaryAmount        EpiMonetaryAmount        `xml:",chardata"`
	AmountCurrencyIdentifier AmountCurrencyIdentifier `xml:"AmountCurrencyIdentifier,attr"`
}

type EpiBeneficiaryPartyDetailsType struct {
	EpiNameAddressDetails GenericTokenType235   `xml:"EpiNameAddressDetails,omitempty"`
	EpiBei                GenericNMtokenType811 `xml:"EpiBei,omitempty"`
	EpiAccountID          EpiAccountIDType      `xml:"EpiAccountID"`
}

type EpiBfiIdentifierType struct {
	GenericNMtokenType811    GenericNMtokenType811    `xml:",chardata"`
	IdentificationSchemeName IdentificationSchemeName `xml:"IdentificationSchemeName,attr"`
}

type EpiBfiPartyDetailsType struct {
	EpiBfiIdentifier EpiBfiIdentifierType `xml:"EpiBfiIdentifier,omitempty"`
	EpiBfiName       GenericStringType135 `xml:"EpiBfiName,omitempty"`
}

type EpiChargeType struct {
	Value        string       `xml:",chardata"`
	ChargeOption ChargeOption `xml:"ChargeOption,attr"`
}

type EpiDetailsType struct {
	EpiIdentificationDetails     EpiIdentificationDetailsType     `xml:"EpiIdentificationDetails"`
	EpiPartyDetails              EpiPartyDetailsType              `xml:"EpiPartyDetails"`
	EpiPaymentInstructionDetails EpiPaymentInstructionDetailsType `xml:"EpiPaymentInstructionDetails"`
}

type EpiIdentificationDetailsType struct {
	EpiDate      Date     `xml:"EpiDate"`
	EpiReference string   `xml:"EpiReference"`
	EpiUrl       string   `xml:"EpiUrl,omitempty"`
	EpiEmail     string   `xml:"EpiEmail,omitempty"`
	EpiOrderInfo []string `xml:"EpiOrderInfo,omitempty"`
}

// Must match the pattern -?[0-9]{1,15},[0-9]{2}
type EpiMonetaryAmount string

type EpiPartyDetailsType struct {
	EpiBfiPartyDetails         EpiBfiPartyDetailsType         `xml:"EpiBfiPartyDetails"`
	EpiBeneficiaryPartyDetails EpiBeneficiaryPartyDetailsType `xml:"EpiBeneficiaryPartyDetails"`
}

type EpiPaymentInstructionDetailsType struct {
	EpiPaymentInstructionId     string                          `xml:"EpiPaymentInstructionId,omitempty"`
	EpiTransactionTypeCode      GenericTokenType3               `xml:"EpiTransactionTypeCode,omitempty"`
	EpiInstructionCode          string                          `xml:"EpiInstructionCode,omitempty"`
	EpiRemittanceInfoIdentifier EpiRemittanceInfoIdentifierType `xml:"EpiRemittanceInfoIdentifier,omitempty"`
	EpiInstructedAmount         EpiAmount                       `xml:"EpiInstructedAmount"`
	EpiCharge                   EpiChargeType                   `xml:"EpiCharge"`
	EpiDateOptionDate           Date                            `xml:"EpiDateOptionDate"`
	EpiPaymentMeansCode         string                          `xml:"EpiPaymentMeansCode,omitempty"`
	EpiPaymentMeansText         GenericStringType170            `xml:"EpiPaymentMeansText,omitempty"`
}

// Must match the pattern ([0-9]{2,20})|(RF[0-9][0-9][0-9A-Za-z]{1,21})
type EpiRemittanceInfoIdentifierPattern string

type EpiRemittanceInfoIdentifierType struct {
	EpiRemittanceInfoIdentifierPattern EpiRemittanceInfoIdentifierPattern `xml:",chardata"`
	IdentificationSchemeName           IdentificationSchemeName           `xml:"IdentificationSchemeName,attr,omitempty"`
}

// Must match the pattern [0-9]{1,15}(,[0-9]{1,6})?
type ExchangeRate string

type ExternalSpecificationDetailsType []string

func (a ExternalSpecificationDetailsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:"item"`
	}
	output.Items = []string(a)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (a *ExternalSpecificationDetailsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				*a = append(*a, item)
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
}

type FactoringAgreementDetailsType struct {
	FactoringAgreementIdentifier       string                             `xml:"FactoringAgreementIdentifier"`
	TransmissionListIdentifier         string                             `xml:"TransmissionListIdentifier,omitempty"`
	EndorsementClauseCode              string                             `xml:"EndorsementClauseCode,omitempty"`
	FactoringTypeCode                  string                             `xml:"FactoringTypeCode,omitempty"`
	FactoringFreeText                  []string                           `xml:"FactoringFreeText,omitempty"`
	FactoringPartyIdentifier           PartyLegalRegIdType                `xml:"FactoringPartyIdentifier,omitempty"`
	FactoringPartyName                 string                             `xml:"FactoringPartyName,omitempty"`
	FactoringPartyPostalAddressDetails FactoringPartyPostalAddressDetails `xml:"FactoringPartyPostalAddressDetails,omitempty"`
}

type FactoringPartyPostalAddressDetails struct {
	FactoringPartyStreetName              []GenericStringType235 `xml:"FactoringPartyStreetName"`
	FactoringPartyTownName                GenericStringType235   `xml:"FactoringPartyTownName"`
	FactoringPartyPostCodeIdentifier      GenericStringType235   `xml:"FactoringPartyPostCodeIdentifier"`
	FactoringPartyCountrySubdivision      GenericStringType235   `xml:"FactoringPartyCountrySubdivision,omitempty"`
	CountryCode                           CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                           string                 `xml:"CountryName,omitempty"`
	FactoringPartyPostOfficeBoxIdentifier string                 `xml:"FactoringPartyPostOfficeBoxIdentifier,omitempty"`
}

type Finvoice struct {
	MessageTransmissionDetails              MessageTransmissionDetailsType           `xml:"MessageTransmissionDetails"`
	SellerPartyDetails                      SellerPartyDetailsType                   `xml:"SellerPartyDetails"`
	SellerOrganisationUnitNumber            string                                   `xml:"SellerOrganisationUnitNumber,omitempty"`
	SellerSiteCode                          string                                   `xml:"SellerSiteCode,omitempty"`
	SellerContactPersonName                 string                                   `xml:"SellerContactPersonName,omitempty"`
	SellerContactPersonFunction             []string                                 `xml:"SellerContactPersonFunction,omitempty"`
	SellerContactPersonDepartment           []string                                 `xml:"SellerContactPersonDepartment,omitempty"`
	SellerCommunicationDetails              SellerCommunicationDetailsType           `xml:"SellerCommunicationDetails,omitempty"`
	SellerInformationDetails                SellerInformationDetailsType             `xml:"SellerInformationDetails,omitempty"`
	InvoiceSenderPartyDetails               InvoiceSenderPartyDetailsType            `xml:"InvoiceSenderPartyDetails,omitempty"`
	InvoiceRecipientPartyDetails            InvoiceRecipientPartyDetailsType         `xml:"InvoiceRecipientPartyDetails,omitempty"`
	InvoiceRecipientOrganisationUnitNumber  string                                   `xml:"InvoiceRecipientOrganisationUnitNumber,omitempty"`
	InvoiceRecipientSiteCode                string                                   `xml:"InvoiceRecipientSiteCode,omitempty"`
	InvoiceRecipientContactPersonName       string                                   `xml:"InvoiceRecipientContactPersonName,omitempty"`
	InvoiceRecipientContactPersonFunction   []string                                 `xml:"InvoiceRecipientContactPersonFunction,omitempty"`
	InvoiceRecipientContactPersonDepartment []string                                 `xml:"InvoiceRecipientContactPersonDepartment,omitempty"`
	InvoiceRecipientLanguageCode            LanguageCodeType                         `xml:"InvoiceRecipientLanguageCode,omitempty"`
	InvoiceRecipientCommunicationDetails    InvoiceRecipientCommunicationDetailsType `xml:"InvoiceRecipientCommunicationDetails,omitempty"`
	BuyerPartyDetails                       BuyerPartyDetailsType                    `xml:"BuyerPartyDetails"`
	BuyerOrganisationUnitNumber             string                                   `xml:"BuyerOrganisationUnitNumber,omitempty"`
	BuyerSiteCode                           string                                   `xml:"BuyerSiteCode,omitempty"`
	BuyerContactPersonName                  string                                   `xml:"BuyerContactPersonName,omitempty"`
	BuyerContactPersonFunction              []string                                 `xml:"BuyerContactPersonFunction,omitempty"`
	BuyerContactPersonDepartment            []string                                 `xml:"BuyerContactPersonDepartment,omitempty"`
	BuyerCommunicationDetails               BuyerCommunicationDetailsType            `xml:"BuyerCommunicationDetails,omitempty"`
	DeliveryPartyDetails                    DeliveryPartyDetailsType                 `xml:"DeliveryPartyDetails,omitempty"`
	DeliveryOrganisationUnitNumber          string                                   `xml:"DeliveryOrganisationUnitNumber,omitempty"`
	DeliverySiteCode                        string                                   `xml:"DeliverySiteCode,omitempty"`
	DeliveryContactPersonName               string                                   `xml:"DeliveryContactPersonName,omitempty"`
	DeliveryContactPersonFunction           []string                                 `xml:"DeliveryContactPersonFunction,omitempty"`
	DeliveryContactPersonDepartment         []string                                 `xml:"DeliveryContactPersonDepartment,omitempty"`
	DeliveryCommunicationDetails            DeliveryCommunicationDetailsType         `xml:"DeliveryCommunicationDetails,omitempty"`
	DeliveryDetails                         DeliveryDetailsType                      `xml:"DeliveryDetails,omitempty"`
	AnyPartyDetails                         []AnyPartyDetailsType                    `xml:"AnyPartyDetails,omitempty"`
	InvoiceDetails                          InvoiceDetailsType                       `xml:"InvoiceDetails"`
	PaymentCardInfo                         PaymentCardInfoType                      `xml:"PaymentCardInfo,omitempty"`
	DirectDebitInfo                         DirectDebitInfoType                      `xml:"DirectDebitInfo,omitempty"`
	PaymentStatusDetails                    PaymentStatusDetailsType                 `xml:"PaymentStatusDetails,omitempty"`
	PartialPaymentDetails                   []PartialPaymentDetailsType              `xml:"PartialPaymentDetails,omitempty"`
	FactoringAgreementDetails               FactoringAgreementDetailsType            `xml:"FactoringAgreementDetails,omitempty"`
	VirtualBankBarcode                      string                                   `xml:"VirtualBankBarcode,omitempty"`
	InvoiceRow                              []InvoiceRowType                         `xml:"InvoiceRow"`
	SpecificationDetails                    SpecificationDetailsType                 `xml:"SpecificationDetails,omitempty"`
	EpiDetails                              EpiDetailsType                           `xml:"EpiDetails"`
	InvoiceUrlNameText                      []string                                 `xml:"InvoiceUrlNameText,omitempty"`
	InvoiceUrlText                          []string                                 `xml:"InvoiceUrlText,omitempty"`
	StorageUrlText                          string                                   `xml:"StorageUrlText,omitempty"`
	LayOutIdentifier                        string                                   `xml:"LayOutIdentifier,omitempty"`
	InvoiceSegmentIdentifier                string                                   `xml:"InvoiceSegmentIdentifier,omitempty"`
	ControlChecksum                         string                                   `xml:"ControlChecksum,omitempty"`
	MessageChecksum                         string                                   `xml:"MessageChecksum,omitempty"`
	ControlStampText                        string                                   `xml:"ControlStampText,omitempty"`
	AcceptanceStampText                     string                                   `xml:"AcceptanceStampText,omitempty"`
	OriginalInvoiceFormat                   string                                   `xml:"OriginalInvoiceFormat,omitempty"`
	AttachmentMessageDetails                AttachmentMessageDetailsType             `xml:"AttachmentMessageDetails,omitempty"`
	Version                                 Version                                  `xml:"Version,attr"`
}

func (f *Finvoice) MessageID() message.ID {
	return message.ID(f.MessageTransmissionDetails.MessageDetails.MessageIdentifier)
}

func (f *Finvoice) MessageType() message.MessageType {
	return message.MessageType("Finvoice")
}

func (f *Finvoice) From() AddressID {
	return f.MessageTransmissionDetails.MessageSenderDetails.FromIdentifier.Address
}

func (f *Finvoice) FromIntermediator() IntermediatorID {
	return f.MessageTransmissionDetails.MessageSenderDetails.FromIntermediator
}

func (f *Finvoice) To() AddressID {
	return f.MessageTransmissionDetails.MessageReceiverDetails.ToIdentifier.Address
}

func (f *Finvoice) ToIntermediator() IntermediatorID {
	return f.MessageTransmissionDetails.MessageReceiverDetails.ToIntermediator
}

// May be one of CCYYMMDD
type Format string

// Must be at least 1 items long
type GenericNMtokenType134 string

// Must be at least 2 items long
type GenericNMtokenType235 string

// Must be at least 8 items long
type GenericNMtokenType811 string

// Must be at least 1 items long
type GenericStringType110 string

// Must be at least 1 items long
type GenericStringType113 string

// Must be at least 1 items long
type GenericStringType119 string

// Must be at least 1 items long
type GenericStringType120 string

// Must be at least 1 items long
type GenericStringType135 string

// Must be at least 1 items long
type GenericStringType14 string

// Must be at least 1 items long
type GenericStringType170 string

// Must be at least 1 items long
type GenericStringType18 string

// Must be at least 2 items long
type GenericStringType235 string

// Must be at least 2 items long
type GenericStringType248 string

// Must be at least 2 items long
type GenericStringType270 string

// Must be at least 1 items long
type GenericTokenType120 string

// Must be at least 2 items long
type GenericTokenType235 string

// May be no more than 3 items long
type GenericTokenType3 string

type HeaderValueType struct {
	Header GenericStringType135   `xml:"Header,omitempty"`
	Value  []GenericStringType170 `xml:"Value,omitempty"`
}

// May be one of IBAN, BBAN
type IdentificationSchemeName string

type InvoiceChargeDetailsType struct {
	ReasonText      GenericStringType170 `xml:"ReasonText,omitempty"`
	ReasonCode      string               `xml:"ReasonCode,omitempty"`
	Percent         Percentage           `xml:"Percent,omitempty"`
	Amount          Amount               `xml:"Amount,omitempty"`
	BaseAmount      Amount               `xml:"BaseAmount,omitempty"`
	VatCategoryCode string               `xml:"VatCategoryCode,omitempty"`
	VatRatePercent  Percentage           `xml:"VatRatePercent,omitempty"`
}

type InvoiceClassificationType struct {
	ClassificationCode GenericStringType110 `xml:"ClassificationCode,omitempty"`
	ClassificationText GenericStringType170 `xml:"ClassificationText,omitempty"`
}

type InvoiceDetailsType struct {
	InvoiceTypeCode                      InvoiceTypeCodeTypeFI          `xml:"InvoiceTypeCode"`
	InvoiceTypeCodeUN                    string                         `xml:"InvoiceTypeCodeUN,omitempty"`
	InvoiceTypeText                      GenericStringType135           `xml:"InvoiceTypeText"`
	InvoiceClassification                InvoiceClassificationType      `xml:"InvoiceClassification,omitempty"`
	OriginCode                           OriginCodeType                 `xml:"OriginCode"`
	OriginText                           string                         `xml:"OriginText,omitempty"`
	InvoicedObjectID                     InvoicedObjectIDType           `xml:"InvoicedObjectID,omitempty"`
	InvoiceNumber                        GenericStringType120           `xml:"InvoiceNumber"`
	InvoiceDate                          Date                           `xml:"InvoiceDate"`
	OriginalInvoiceNumber                GenericStringType120           `xml:"OriginalInvoiceNumber,omitempty"`
	OriginalInvoiceDate                  Date                           `xml:"OriginalInvoiceDate,omitempty"`
	OriginalInvoiceReference             []OriginalInvoiceReferenceType `xml:"OriginalInvoiceReference,omitempty"`
	InvoicingPeriodStartDate             Date                           `xml:"InvoicingPeriodStartDate,omitempty"`
	InvoicingPeriodEndDate               Date                           `xml:"InvoicingPeriodEndDate,omitempty"`
	SellerReferenceIdentifier            string                         `xml:"SellerReferenceIdentifier,omitempty"`
	SellerReferenceIdentifierUrlText     string                         `xml:"SellerReferenceIdentifierUrlText,omitempty"`
	BuyersSellerIdentifier               PartyIdentifierType            `xml:"BuyersSellerIdentifier,omitempty"`
	SellersBuyerIdentifier               PartyIdentifierType            `xml:"SellersBuyerIdentifier,omitempty"`
	OrderIdentifier                      string                         `xml:"OrderIdentifier,omitempty"`
	OrderIdentifierUrlText               string                         `xml:"OrderIdentifierUrlText,omitempty"`
	OrderDate                            Date                           `xml:"OrderDate,omitempty"`
	OrdererName                          string                         `xml:"OrdererName,omitempty"`
	SalesPersonName                      string                         `xml:"SalesPersonName,omitempty"`
	OrderConfirmationIdentifier          string                         `xml:"OrderConfirmationIdentifier,omitempty"`
	OrderConfirmationDate                Date                           `xml:"OrderConfirmationDate,omitempty"`
	AgreementIdentifier                  string                         `xml:"AgreementIdentifier,omitempty"`
	AgreementIdentifierUrlText           string                         `xml:"AgreementIdentifierUrlText,omitempty"`
	AgreementTypeText                    string                         `xml:"AgreementTypeText,omitempty"`
	AgreementTypeCode                    string                         `xml:"AgreementTypeCode,omitempty"`
	AgreementDate                        Date                           `xml:"AgreementDate,omitempty"`
	NotificationIdentifier               string                         `xml:"NotificationIdentifier,omitempty"`
	NotificationDate                     Date                           `xml:"NotificationDate,omitempty"`
	RegistrationNumberIdentifier         string                         `xml:"RegistrationNumberIdentifier,omitempty"`
	ControllerIdentifier                 string                         `xml:"ControllerIdentifier,omitempty"`
	ControllerName                       string                         `xml:"ControllerName,omitempty"`
	ControlDate                          Date                           `xml:"ControlDate,omitempty"`
	BuyerReferenceIdentifier             string                         `xml:"BuyerReferenceIdentifier,omitempty"`
	ProjectReferenceIdentifier           string                         `xml:"ProjectReferenceIdentifier,omitempty"`
	DefinitionDetails                    []DefinitionDetails            `xml:"DefinitionDetails,omitempty"`
	RowsTotalVatExcludedAmount           Amount                         `xml:"RowsTotalVatExcludedAmount,omitempty"`
	DiscountsTotalVatExcludedAmount      Amount                         `xml:"DiscountsTotalVatExcludedAmount,omitempty"`
	ChargesTotalVatExcludedAmount        Amount                         `xml:"ChargesTotalVatExcludedAmount,omitempty"`
	InvoiceTotalVatExcludedAmount        Amount                         `xml:"InvoiceTotalVatExcludedAmount,omitempty"`
	InvoiceTotalVatAmount                Amount                         `xml:"InvoiceTotalVatAmount,omitempty"`
	InvoiceTotalVatAccountingAmount      Amount                         `xml:"InvoiceTotalVatAccountingAmount,omitempty"`
	InvoiceTotalVatIncludedAmount        Amount                         `xml:"InvoiceTotalVatIncludedAmount"`
	InvoiceTotalRoundoffAmount           Amount                         `xml:"InvoiceTotalRoundoffAmount,omitempty"`
	InvoicePaidAmount                    Amount                         `xml:"InvoicePaidAmount,omitempty"`
	ExchangeRate                         ExchangeRate                   `xml:"ExchangeRate,omitempty"`
	OtherCurrencyAmountVatExcludedAmount Amount                         `xml:"OtherCurrencyAmountVatExcludedAmount,omitempty"`
	OtherCurrencyAmountVatIncludedAmount Amount                         `xml:"OtherCurrencyAmountVatIncludedAmount,omitempty"`
	CreditLimitAmount                    Amount                         `xml:"CreditLimitAmount,omitempty"`
	CreditInterestPercent                Percentage                     `xml:"CreditInterestPercent,omitempty"`
	OperationLimitAmount                 Amount                         `xml:"OperationLimitAmount,omitempty"`
	MonthlyAmount                        Amount                         `xml:"MonthlyAmount,omitempty"`
	ShortProposedAccountIdentifier       string                         `xml:"ShortProposedAccountIdentifier,omitempty"`
	NormalProposedAccountIdentifier      string                         `xml:"NormalProposedAccountIdentifier,omitempty"`
	ProposedAccountText                  string                         `xml:"ProposedAccountText,omitempty"`
	AccountDimensionText                 string                         `xml:"AccountDimensionText,omitempty"`
	SellerAccountText                    string                         `xml:"SellerAccountText,omitempty"`
	VatPoint                             VatPointType                   `xml:"VatPoint,omitempty"`
	VatSpecificationDetails              []VatSpecificationDetailsType  `xml:"VatSpecificationDetails,omitempty"`
	InvoiceFreeText                      []string                       `xml:"InvoiceFreeText,omitempty"`
	InvoiceVatFreeText                   string                         `xml:"InvoiceVatFreeText,omitempty"`
	PaymentTermsDetails                  []PaymentTermsDetailsType      `xml:"PaymentTermsDetails,omitempty"`
	DiscountDetails                      []DiscountDetailsType          `xml:"DiscountDetails,omitempty"`
	ChargeDetails                        []InvoiceChargeDetailsType     `xml:"ChargeDetails,omitempty"`
	TenderReference                      GenericStringType170           `xml:"TenderReference,omitempty"`
}

type InvoiceRecipientCommunicationDetailsType struct {
	InvoiceRecipientPhoneNumberIdentifier  string `xml:"InvoiceRecipientPhoneNumberIdentifier,omitempty"`
	InvoiceRecipientEmailaddressIdentifier string `xml:"InvoiceRecipientEmailaddressIdentifier,omitempty"`
}

type InvoiceRecipientDetailsType struct {
	InvoiceRecipientAddress              GenericStringType135  `xml:"InvoiceRecipientAddress"`
	InvoiceRecipientIntermediatorAddress GenericNMtokenType811 `xml:"InvoiceRecipientIntermediatorAddress"`
}

type InvoiceRecipientPartyDetailsType struct {
	InvoiceRecipientPartyIdentifier      PartyLegalRegIdType                      `xml:"InvoiceRecipientPartyIdentifier,omitempty"`
	InvoiceRecipientOrganisationName     []GenericStringType235                   `xml:"InvoiceRecipientOrganisationName"`
	InvoiceRecipientDepartment           []string                                 `xml:"InvoiceRecipientDepartment,omitempty"`
	InvoiceRecipientOrganisationTaxCode  string                                   `xml:"InvoiceRecipientOrganisationTaxCode,omitempty"`
	InvoiceRecipientCode                 PartyIdentifierType                      `xml:"InvoiceRecipientCode,omitempty"`
	InvoiceRecipientPostalAddressDetails InvoiceRecipientPostalAddressDetailsType `xml:"InvoiceRecipientPostalAddressDetails,omitempty"`
}

type InvoiceRecipientPostalAddressDetailsType struct {
	InvoiceRecipientStreetName              []GenericStringType235 `xml:"InvoiceRecipientStreetName"`
	InvoiceRecipientTownName                GenericStringType235   `xml:"InvoiceRecipientTownName"`
	InvoiceRecipientPostCodeIdentifier      GenericStringType235   `xml:"InvoiceRecipientPostCodeIdentifier"`
	InvoiceRecipientCountrySubdivision      GenericStringType235   `xml:"InvoiceRecipientCountrySubdivision,omitempty"`
	CountryCode                             CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                             string                 `xml:"CountryName,omitempty"`
	InvoiceRecipientPostOfficeBoxIdentifier string                 `xml:"InvoiceRecipientPostOfficeBoxIdentifier,omitempty"`
}

type InvoiceRowType struct {
	RowSubIdentifier                       string                              `xml:"RowSubIdentifier,omitempty"`
	InvoicedObjectID                       InvoicedObjectIDType                `xml:"InvoicedObjectID,omitempty"`
	ArticleIdentifier                      string                              `xml:"ArticleIdentifier,omitempty"`
	ArticleGroupIdentifier                 []ArticleGroupIdentifierType        `xml:"ArticleGroupIdentifier,omitempty"`
	ArticleName                            string                              `xml:"ArticleName,omitempty"`
	ArticleDescription                     string                              `xml:"ArticleDescription,omitempty"`
	ArticleInfoUrlText                     string                              `xml:"ArticleInfoUrlText,omitempty"`
	BuyerArticleIdentifier                 string                              `xml:"BuyerArticleIdentifier,omitempty"`
	EanCode                                EanCodeType                         `xml:"EanCode,omitempty"`
	RowRegistrationNumberIdentifier        string                              `xml:"RowRegistrationNumberIdentifier,omitempty"`
	SerialNumberIdentifier                 string                              `xml:"SerialNumberIdentifier,omitempty"`
	RowActionCode                          string                              `xml:"RowActionCode,omitempty"`
	RowDefinitionDetails                   []RowDefinitionDetailsType          `xml:"RowDefinitionDetails,omitempty"`
	OfferedQuantity                        []QuantityType014                   `xml:"OfferedQuantity,omitempty"`
	DeliveredQuantity                      []QuantityType014                   `xml:"DeliveredQuantity,omitempty"`
	OrderedQuantity                        QuantityType014                     `xml:"OrderedQuantity,omitempty"`
	ConfirmedQuantity                      QuantityType014                     `xml:"ConfirmedQuantity,omitempty"`
	PostDeliveredQuantity                  QuantityType014                     `xml:"PostDeliveredQuantity,omitempty"`
	InvoicedQuantity                       []QuantityType014                   `xml:"InvoicedQuantity,omitempty"`
	CreditRequestedQuantity                QuantityType014                     `xml:"CreditRequestedQuantity,omitempty"`
	ReturnedQuantity                       QuantityType014                     `xml:"ReturnedQuantity,omitempty"`
	StartDate                              Date                                `xml:"StartDate,omitempty"`
	EndDate                                Date                                `xml:"EndDate,omitempty"`
	UnitPriceAmount                        UnitAmountUN                        `xml:"UnitPriceAmount,omitempty"`
	UnitPriceDiscountAmount                UnitAmountUN                        `xml:"UnitPriceDiscountAmount,omitempty"`
	UnitPriceNetAmount                     UnitAmountUN                        `xml:"UnitPriceNetAmount,omitempty"`
	UnitPriceVatIncludedAmount             UnitAmountUN                        `xml:"UnitPriceVatIncludedAmount,omitempty"`
	UnitPriceBaseQuantity                  QuantityType014                     `xml:"UnitPriceBaseQuantity,omitempty"`
	RowIdentifier                          string                              `xml:"RowIdentifier,omitempty"`
	RowIdentifierUrlText                   string                              `xml:"RowIdentifierUrlText,omitempty"`
	RowOrderPositionIdentifier             string                              `xml:"RowOrderPositionIdentifier,omitempty"`
	RowIdentifierDate                      Date                                `xml:"RowIdentifierDate,omitempty"`
	RowPositionIdentifier                  string                              `xml:"RowPositionIdentifier,omitempty"`
	OriginalInvoiceNumber                  GenericStringType120                `xml:"OriginalInvoiceNumber,omitempty"`
	OriginalInvoiceDate                    Date                                `xml:"OriginalInvoiceDate,omitempty"`
	OriginalInvoiceReference               []OriginalInvoiceReferenceType      `xml:"OriginalInvoiceReference,omitempty"`
	RowOrdererName                         string                              `xml:"RowOrdererName,omitempty"`
	RowSalesPersonName                     string                              `xml:"RowSalesPersonName,omitempty"`
	RowOrderConfirmationIdentifier         string                              `xml:"RowOrderConfirmationIdentifier,omitempty"`
	RowOrderConfirmationDate               Date                                `xml:"RowOrderConfirmationDate,omitempty"`
	RowDeliveryIdentifier                  string                              `xml:"RowDeliveryIdentifier,omitempty"`
	RowDeliveryIdentifierUrlText           string                              `xml:"RowDeliveryIdentifierUrlText,omitempty"`
	RowDeliveryDate                        Date                                `xml:"RowDeliveryDate,omitempty"`
	RowQuotationIdentifier                 string                              `xml:"RowQuotationIdentifier,omitempty"`
	RowQuotationIdentifierUrlText          string                              `xml:"RowQuotationIdentifierUrlText,omitempty"`
	RowAgreementIdentifier                 string                              `xml:"RowAgreementIdentifier,omitempty"`
	RowAgreementIdentifierUrlText          string                              `xml:"RowAgreementIdentifierUrlText,omitempty"`
	RowRequestOfQuotationIdentifier        string                              `xml:"RowRequestOfQuotationIdentifier,omitempty"`
	RowRequestOfQuotationIdentifierUrlText string                              `xml:"RowRequestOfQuotationIdentifierUrlText,omitempty"`
	RowPriceListIdentifier                 string                              `xml:"RowPriceListIdentifier,omitempty"`
	RowPriceListIdentifierUrlText          string                              `xml:"RowPriceListIdentifierUrlText,omitempty"`
	RowBuyerReferenceIdentifier            string                              `xml:"RowBuyerReferenceIdentifier,omitempty"`
	RowProjectReferenceIdentifier          string                              `xml:"RowProjectReferenceIdentifier,omitempty"`
	RowOverDuePaymentDetails               RowOverDuePaymentDetailsType        `xml:"RowOverDuePaymentDetails,omitempty"`
	RowAnyPartyDetails                     []RowAnyPartyDetailsType            `xml:"RowAnyPartyDetails,omitempty"`
	RowDeliveryDetails                     RowDeliveryDetailsType              `xml:"RowDeliveryDetails,omitempty"`
	RowShortProposedAccountIdentifier      string                              `xml:"RowShortProposedAccountIdentifier,omitempty"`
	RowNormalProposedAccountIdentifier     string                              `xml:"RowNormalProposedAccountIdentifier,omitempty"`
	RowProposedAccountText                 string                              `xml:"RowProposedAccountText,omitempty"`
	RowAccountDimensionText                string                              `xml:"RowAccountDimensionText,omitempty"`
	RowSellerAccountText                   string                              `xml:"RowSellerAccountText,omitempty"`
	RowFreeText                            []string                            `xml:"RowFreeText,omitempty"`
	RowUsedQuantity                        QuantityType014                     `xml:"RowUsedQuantity,omitempty"`
	RowPreviousMeterReadingDate            Date                                `xml:"RowPreviousMeterReadingDate,omitempty"`
	RowLatestMeterReadingDate              Date                                `xml:"RowLatestMeterReadingDate,omitempty"`
	RowCalculatedQuantity                  QuantityType014                     `xml:"RowCalculatedQuantity,omitempty"`
	RowAveragePriceAmount                  Amount                              `xml:"RowAveragePriceAmount,omitempty"`
	RowDiscountPercent                     Percentage                          `xml:"RowDiscountPercent,omitempty"`
	RowDiscountAmount                      Amount                              `xml:"RowDiscountAmount,omitempty"`
	RowDiscountBaseAmount                  Amount                              `xml:"RowDiscountBaseAmount,omitempty"`
	RowDiscountTypeCode                    string                              `xml:"RowDiscountTypeCode,omitempty"`
	RowDiscountTypeText                    string                              `xml:"RowDiscountTypeText,omitempty"`
	RowProgressiveDiscountDetails          []RowProgressiveDiscountDetailsType `xml:"RowProgressiveDiscountDetails,omitempty"`
	RowChargeDetails                       []RowChargeDetailsType              `xml:"RowChargeDetails,omitempty"`
	RowVatRatePercent                      Percentage                          `xml:"RowVatRatePercent,omitempty"`
	RowVatCode                             string                              `xml:"RowVatCode,omitempty"`
	RowVatAmount                           Amount                              `xml:"RowVatAmount,omitempty"`
	RowVatExcludedAmount                   Amount                              `xml:"RowVatExcludedAmount,omitempty"`
	RowAmount                              Amount                              `xml:"RowAmount,omitempty"`
	RowTransactionDetails                  TransactionDetailsType              `xml:"RowTransactionDetails,omitempty"`
	SubInvoiceRow                          []SubInvoiceRowType                 `xml:"SubInvoiceRow,omitempty"`
}

type InvoiceSenderPartyDetailsType struct {
	InvoiceSenderPartyIdentifier     PartyLegalRegIdType    `xml:"InvoiceSenderPartyIdentifier,omitempty"`
	InvoiceSenderOrganisationName    []GenericStringType235 `xml:"InvoiceSenderOrganisationName"`
	InvoiceSenderOrganisationTaxCode string                 `xml:"InvoiceSenderOrganisationTaxCode,omitempty"`
	InvoiceSenderCode                PartyIdentifierType    `xml:"InvoiceSenderCode,omitempty"`
}

// Must match the pattern (REQ|QUO|ORD|ORC|INV|DEV|TES|INF|PRI|DEN|SEI|REC|RES|SDD)[0-9]{2}
type InvoiceTypeCodePatternFI string

type InvoiceTypeCodeTypeFI struct {
	InvoiceTypeCodePatternFI InvoiceTypeCodePatternFI `xml:",chardata"`
	CodeListAgencyIdentifier CodeListAgencyIdentifier `xml:"CodeListAgencyIdentifier,attr,omitempty"`
}

type InvoicedObjectIDType struct {
	GenericStringType170 GenericStringType170 `xml:",chardata"`
	SchemeID             string               `xml:"SchemeID,attr,omitempty"`
}

// Must match the pattern [0-9]{4}
type Iso6523cid string

// May be no more than 2 items long
type LanguageCodeType string

type MessageDetails struct {
	MessageIdentifier       GenericStringType248 `xml:"MessageIdentifier"`
	MessageTimeStamp        time.Time            `xml:"MessageTimeStamp"`
	RefToMessageIdentifier  string               `xml:"RefToMessageIdentifier,omitempty"`
	ImplementationCode      string               `xml:"ImplementationCode,omitempty"`
	SpecificationIdentifier GenericStringType135 `xml:"SpecificationIdentifier,omitempty"`
}

func (t *MessageDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MessageDetails
	var layout struct {
		*T
		MessageTimeStamp *xsdDateTime `xml:"MessageTimeStamp"`
	}
	layout.T = (*T)(t)
	layout.MessageTimeStamp = (*xsdDateTime)(&layout.T.MessageTimeStamp)
	return e.EncodeElement(layout, start)
}
func (t *MessageDetails) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MessageDetails
	var overlay struct {
		*T
		MessageTimeStamp *xsdDateTime `xml:"MessageTimeStamp"`
	}
	overlay.T = (*T)(t)
	overlay.MessageTimeStamp = (*xsdDateTime)(&overlay.T.MessageTimeStamp)
	return d.DecodeElement(&overlay, &start)
}

// May be no less that 2 character or no more that 35 characters
type IntermediatorID string

type MessageReceiverDetails struct {
	ToIdentifier    FinvoiceAddress `xml:"ToIdentifier"`
	ToIntermediator IntermediatorID `xml:"ToIntermediator"`
}

type MessageSenderDetails struct {
	FromIdentifier    FinvoiceAddress `xml:"FromIdentifier"`
	FromIntermediator IntermediatorID `xml:"FromIntermediator"`
}

type MessageTransmissionDetailsType struct {
	MessageSenderDetails   MessageSenderDetails   `xml:"MessageSenderDetails"`
	MessageReceiverDetails MessageReceiverDetails `xml:"MessageReceiverDetails"`
	MessageDetails         MessageDetails         `xml:"MessageDetails"`
}

// Must match the pattern -?[0-9]{1,15}(,[0-9]{2,5})?
type MonetaryAmount string

// May be one of Original, Copy, Cancel
type OriginCodeType string

type OriginalInvoiceReferenceType struct {
	InvoiceNumber GenericStringType120 `xml:"InvoiceNumber,omitempty"`
	InvoiceDate   Date                 `xml:"InvoiceDate,omitempty"`
}

type PackageDetails struct {
	PackageLength             QuantityType014 `xml:"PackageLength,omitempty"`
	PackageWidth              QuantityType014 `xml:"PackageWidth,omitempty"`
	PackageHeight             QuantityType014 `xml:"PackageHeight,omitempty"`
	PackageWeight             QuantityType014 `xml:"PackageWeight,omitempty"`
	PackageNetWeight          QuantityType014 `xml:"PackageNetWeight,omitempty"`
	PackageVolume             QuantityType014 `xml:"PackageVolume,omitempty"`
	TransportCarriageQuantity QuantityType014 `xml:"TransportCarriageQuantity,omitempty"`
}

type PartialPaymentDetailsType struct {
	PaidAmount                        Amount                 `xml:"PaidAmount"`
	PaidVatExcludedAmount             Amount                 `xml:"PaidVatExcludedAmount,omitempty"`
	UnPaidAmount                      Amount                 `xml:"UnPaidAmount"`
	UnPaidVatExcludedAmount           Amount                 `xml:"UnPaidVatExcludedAmount,omitempty"`
	InterestPercent                   Percentage             `xml:"InterestPercent,omitempty"`
	ProsessingCostsAmount             Amount                 `xml:"ProsessingCostsAmount,omitempty"`
	PartialPaymentVatIncludedAmount   []Amount               `xml:"PartialPaymentVatIncludedAmount"`
	PartialPaymentVatExcludedAmount   []Amount               `xml:"PartialPaymentVatExcludedAmount"`
	PartialPaymentDueDate             []Date                 `xml:"PartialPaymentDueDate"`
	PartialPaymentReferenceIdentifier []GenericStringType235 `xml:"PartialPaymentReferenceIdentifier"`
}

type PartyIdentifierType struct {
	Value          string              `xml:",chardata"`
	IdentifierType GenericTokenType120 `xml:"IdentifierType,attr,omitempty"`
	SchemeID       Iso6523cid          `xml:"SchemeID,attr,omitempty"`
}

type PartyLegalRegIdType struct {
	Value    string     `xml:",chardata"`
	SchemeID Iso6523cid `xml:"SchemeID,attr,omitempty"`
}

type PaymentCardInfoType struct {
	PrimaryAccountNumber GenericStringType119 `xml:"PrimaryAccountNumber"`
	CardHolderName       GenericStringType170 `xml:"CardHolderName,omitempty"`
}

type PaymentOverDueFineDetailsType struct {
	PaymentOverDueFineFreeText []string   `xml:"PaymentOverDueFineFreeText,omitempty"`
	PaymentOverDueFinePercent  Percentage `xml:"PaymentOverDueFinePercent,omitempty"`
	PaymentOverDueFixedAmount  Amount     `xml:"PaymentOverDueFixedAmount,omitempty"`
}

// May be one of PAID, NOTPAID, PARTLYPAID
type PaymentStatusCodeType string

type PaymentStatusDetailsType struct {
	PaymentStatusCode PaymentStatusCodeType `xml:"PaymentStatusCode,omitempty"`
	PaymentMethodText string                `xml:"PaymentMethodText,omitempty"`
}

type PaymentTermsDetailsType struct {
	PaymentTermsFreeText            []string                      `xml:"PaymentTermsFreeText,omitempty"`
	FreeText                        []HeaderValueType             `xml:"FreeText,omitempty"`
	InvoiceDueDate                  Date                          `xml:"InvoiceDueDate,omitempty"`
	CashDiscountDate                Date                          `xml:"CashDiscountDate,omitempty"`
	CashDiscountBaseAmount          Amount                        `xml:"CashDiscountBaseAmount,omitempty"`
	CashDiscountPercent             Percentage                    `xml:"CashDiscountPercent,omitempty"`
	CashDiscountAmount              Amount                        `xml:"CashDiscountAmount,omitempty"`
	CashDiscountExcludingVatAmount  Amount                        `xml:"CashDiscountExcludingVatAmount,omitempty"`
	CashDiscountVatDetails          []CashDiscountVatDetails      `xml:"CashDiscountVatDetails,omitempty"`
	ReducedInvoiceVatIncludedAmount Amount                        `xml:"ReducedInvoiceVatIncludedAmount,omitempty"`
	PaymentOverDueFineDetails       PaymentOverDueFineDetailsType `xml:"PaymentOverDueFineDetails,omitempty"`
}

// Must match the pattern [1-9]?[0-9]{1,2}(,[0-9]{1,4})?
type Percentage string

type QuantityType struct {
	Value              string     `xml:",chardata"`
	QuantityUnitCode   string     `xml:"QuantityUnitCode,attr,omitempty"`
	QuantityUnitCodeUN UnitCodeUN `xml:"QuantityUnitCodeUN,attr,omitempty"`
}

type QuantityType014 struct {
	Value              string     `xml:",chardata"`
	QuantityUnitCode   string     `xml:"QuantityUnitCode,attr,omitempty"`
	QuantityUnitCodeUN UnitCodeUN `xml:"QuantityUnitCodeUN,attr,omitempty"`
}

type QuantityType070 struct {
	Value              string     `xml:",chardata"`
	QuantityUnitCode   string     `xml:"QuantityUnitCode,attr,omitempty"`
	QuantityUnitCodeUN UnitCodeUN `xml:"QuantityUnitCodeUN,attr,omitempty"`
}

type RowAnyPartyDetailsType struct {
	RowAnyPartyText                   Anypartytexttype035             `xml:"RowAnyPartyText"`
	RowAnyPartyIdentifier             PartyLegalRegIdType             `xml:"RowAnyPartyIdentifier,omitempty"`
	RowAnyPartyOrganisationName       []GenericStringType235          `xml:"RowAnyPartyOrganisationName"`
	RowAnyPartyOrganisationDepartment []string                        `xml:"RowAnyPartyOrganisationDepartment,omitempty"`
	RowAnyPartyOrganisationTaxCode    string                          `xml:"RowAnyPartyOrganisationTaxCode,omitempty"`
	RowAnyPartyCode                   PartyIdentifierType             `xml:"RowAnyPartyCode,omitempty"`
	RowAnyPartyPostalAddressDetails   RowAnyPartyPostalAddressDetails `xml:"RowAnyPartyPostalAddressDetails,omitempty"`
	RowAnyPartyOrganisationUnitNumber string                          `xml:"RowAnyPartyOrganisationUnitNumber,omitempty"`
	RowAnyPartySiteCode               string                          `xml:"RowAnyPartySiteCode,omitempty"`
}

type RowAnyPartyPostalAddressDetails struct {
	RowAnyPartyStreetName              []GenericStringType235 `xml:"RowAnyPartyStreetName"`
	RowAnyPartyTownName                GenericStringType235   `xml:"RowAnyPartyTownName"`
	RowAnyPartyPostCodeIdentifier      GenericStringType235   `xml:"RowAnyPartyPostCodeIdentifier"`
	RowAnyPartyCountrySubdivision      GenericStringType235   `xml:"RowAnyPartyCountrySubdivision,omitempty"`
	CountryCode                        CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                        string                 `xml:"CountryName,omitempty"`
	RowAnyPartyPostOfficeBoxIdentifier string                 `xml:"RowAnyPartyPostOfficeBoxIdentifier,omitempty"`
}

type RowChargeDetailsType struct {
	ReasonText GenericStringType170 `xml:"ReasonText,omitempty"`
	ReasonCode string               `xml:"ReasonCode,omitempty"`
	Percent    Percentage           `xml:"Percent,omitempty"`
	Amount     Amount               `xml:"Amount,omitempty"`
	BaseAmount Amount               `xml:"BaseAmount,omitempty"`
}

type RowDefinitionDetailsType struct {
	RowDefinitionHeaderText RowDefinitionHeaderText `xml:"RowDefinitionHeaderText"`
	RowDefinitionValue      QuantityType070         `xml:"RowDefinitionValue,omitempty"`
}

type RowDefinitionHeaderText struct {
	Value          string              `xml:",chardata"`
	DefinitionCode GenericTokenType120 `xml:"DefinitionCode,attr,omitempty"`
}

type RowDeliveryDetailsType struct {
	RowTerminalAddressText            string               `xml:"RowTerminalAddressText,omitempty"`
	RowWaybillIdentifier              string               `xml:"RowWaybillIdentifier,omitempty"`
	RowWaybillTypeCode                string               `xml:"RowWaybillTypeCode,omitempty"`
	RowClearanceIdentifier            string               `xml:"RowClearanceIdentifier,omitempty"`
	RowDeliveryNoteIdentifier         string               `xml:"RowDeliveryNoteIdentifier,omitempty"`
	RowDelivererIdentifier            string               `xml:"RowDelivererIdentifier,omitempty"`
	RowDelivererName                  []string             `xml:"RowDelivererName,omitempty"`
	RowDelivererCountrySubdivision    GenericStringType235 `xml:"RowDelivererCountrySubdivision,omitempty"`
	RowDelivererCountryCode           CountryCodeType      `xml:"RowDelivererCountryCode,omitempty"`
	RowDelivererCountryName           string               `xml:"RowDelivererCountryName,omitempty"`
	RowModeOfTransportIdentifier      string               `xml:"RowModeOfTransportIdentifier,omitempty"`
	RowCarrierName                    string               `xml:"RowCarrierName,omitempty"`
	RowVesselName                     string               `xml:"RowVesselName,omitempty"`
	RowLocationIdentifier             string               `xml:"RowLocationIdentifier,omitempty"`
	RowTransportInformationDate       Date                 `xml:"RowTransportInformationDate,omitempty"`
	RowCountryOfOrigin                string               `xml:"RowCountryOfOrigin,omitempty"`
	RowCountryOfDestinationName       string               `xml:"RowCountryOfDestinationName,omitempty"`
	RowDestinationCountryCode         CountryCodeType      `xml:"RowDestinationCountryCode,omitempty"`
	RowPlaceOfDischarge               string               `xml:"RowPlaceOfDischarge,omitempty"`
	RowFinalDestinationName           []string             `xml:"RowFinalDestinationName,omitempty"`
	RowCustomsInfo                    CustomsInfoType      `xml:"RowCustomsInfo,omitempty"`
	RowManufacturerArticleIdentifier  string               `xml:"RowManufacturerArticleIdentifier,omitempty"`
	RowManufacturerIdentifier         string               `xml:"RowManufacturerIdentifier,omitempty"`
	RowManufacturerName               []string             `xml:"RowManufacturerName,omitempty"`
	RowManufacturerCountrySubdivision GenericStringType235 `xml:"RowManufacturerCountrySubdivision,omitempty"`
	RowManufacturerCountryCode        CountryCodeType      `xml:"RowManufacturerCountryCode,omitempty"`
	RowManufacturerCountryName        string               `xml:"RowManufacturerCountryName,omitempty"`
	RowManufacturerOrderIdentifier    string               `xml:"RowManufacturerOrderIdentifier,omitempty"`
	RowPackageDetails                 RowPackageDetails    `xml:"RowPackageDetails,omitempty"`
}

type RowOverDuePaymentDetailsType struct {
	RowOriginalInvoiceIdentifier           string          `xml:"RowOriginalInvoiceIdentifier,omitempty"`
	RowOriginalInvoiceDate                 Date            `xml:"RowOriginalInvoiceDate,omitempty"`
	RowOriginalDueDate                     Date            `xml:"RowOriginalDueDate,omitempty"`
	RowOriginalInvoiceTotalAmount          Amount          `xml:"RowOriginalInvoiceTotalAmount,omitempty"`
	RowOriginalEpiRemittanceInfoIdentifier string          `xml:"RowOriginalEpiRemittanceInfoIdentifier,omitempty"`
	RowPaidVatExcludedAmount               Amount          `xml:"RowPaidVatExcludedAmount,omitempty"`
	RowPaidVatIncludedAmount               Amount          `xml:"RowPaidVatIncludedAmount,omitempty"`
	RowPaidDate                            Date            `xml:"RowPaidDate,omitempty"`
	RowUnPaidVatExcludedAmount             Amount          `xml:"RowUnPaidVatExcludedAmount,omitempty"`
	RowUnPaidVatIncludedAmount             Amount          `xml:"RowUnPaidVatIncludedAmount,omitempty"`
	RowCollectionDate                      Date            `xml:"RowCollectionDate,omitempty"`
	RowCollectionQuantity                  QuantityType014 `xml:"RowCollectionQuantity,omitempty"`
	RowCollectionChargeAmount              Amount          `xml:"RowCollectionChargeAmount,omitempty"`
	RowInterestRate                        Percentage      `xml:"RowInterestRate,omitempty"`
	RowInterestStartDate                   Date            `xml:"RowInterestStartDate,omitempty"`
	RowInterestEndDate                     Date            `xml:"RowInterestEndDate,omitempty"`
	RowInterestPeriodText                  string          `xml:"RowInterestPeriodText,omitempty"`
	RowInterestDateNumber                  string          `xml:"RowInterestDateNumber,omitempty"`
	RowInterestChargeAmount                Amount          `xml:"RowInterestChargeAmount,omitempty"`
	RowInterestChargeVatAmount             Amount          `xml:"RowInterestChargeVatAmount,omitempty"`
}

type RowPackageDetails struct {
	RowPackageLength             QuantityType014 `xml:"RowPackageLength,omitempty"`
	RowPackageWidth              QuantityType014 `xml:"RowPackageWidth,omitempty"`
	RowPackageHeight             QuantityType014 `xml:"RowPackageHeight,omitempty"`
	RowPackageWeight             QuantityType014 `xml:"RowPackageWeight,omitempty"`
	RowPackageNetWeight          QuantityType014 `xml:"RowPackageNetWeight,omitempty"`
	RowPackageVolume             QuantityType014 `xml:"RowPackageVolume,omitempty"`
	RowTransportCarriageQuantity QuantityType014 `xml:"RowTransportCarriageQuantity,omitempty"`
}

type RowProgressiveDiscountDetailsType struct {
	RowDiscountPercent    Percentage `xml:"RowDiscountPercent,omitempty"`
	RowDiscountAmount     Amount     `xml:"RowDiscountAmount,omitempty"`
	RowDiscountBaseAmount Amount     `xml:"RowDiscountBaseAmount,omitempty"`
	RowDiscountTypeCode   string     `xml:"RowDiscountTypeCode,omitempty"`
	RowDiscountTypeText   string     `xml:"RowDiscountTypeText,omitempty"`
}

type SellerAccountDetailsType struct {
	SellerAccountID   SellerAccountIDType  `xml:"SellerAccountID"`
	SellerBic         SellerBicType        `xml:"SellerBic"`
	SellerAccountName GenericStringType170 `xml:"SellerAccountName,omitempty"`
}

type SellerAccountIDType struct {
	GenericNMtokenType235    GenericNMtokenType235    `xml:",chardata"`
	IdentificationSchemeName IdentificationSchemeName `xml:"IdentificationSchemeName,attr"`
}

type SellerBicType struct {
	GenericNMtokenType811    GenericNMtokenType811    `xml:",chardata"`
	IdentificationSchemeName IdentificationSchemeName `xml:"IdentificationSchemeName,attr"`
}

type SellerCommunicationDetailsType struct {
	SellerPhoneNumberIdentifier  string `xml:"SellerPhoneNumberIdentifier,omitempty"`
	SellerEmailaddressIdentifier string `xml:"SellerEmailaddressIdentifier,omitempty"`
}

type SellerInformationDetailsType struct {
	SellerOfficialPostalAddressDetails SellerOfficialPostalAddressDetails `xml:"SellerOfficialPostalAddressDetails,omitempty"`
	SellerHomeTownName                 string                             `xml:"SellerHomeTownName,omitempty"`
	SellerVatRegistrationText          string                             `xml:"SellerVatRegistrationText,omitempty"`
	SellerVatRegistrationDate          Date                               `xml:"SellerVatRegistrationDate,omitempty"`
	SellerTaxRegistrationText          string                             `xml:"SellerTaxRegistrationText,omitempty"`
	SellerAdditionalLegalInfo          string                             `xml:"SellerAdditionalLegalInfo,omitempty"`
	SellerPhoneNumber                  string                             `xml:"SellerPhoneNumber,omitempty"`
	SellerFaxNumber                    string                             `xml:"SellerFaxNumber,omitempty"`
	SellerCommonEmailaddressIdentifier string                             `xml:"SellerCommonEmailaddressIdentifier,omitempty"`
	SellerWebaddressIdentifier         string                             `xml:"SellerWebaddressIdentifier,omitempty"`
	SellerFreeText                     string                             `xml:"SellerFreeText,omitempty"`
	SellerAccountDetails               []SellerAccountDetailsType         `xml:"SellerAccountDetails,omitempty"`
	InvoiceRecipientDetails            []InvoiceRecipientDetailsType      `xml:"InvoiceRecipientDetails,omitempty"`
}

type SellerOfficialPostalAddressDetails struct {
	SellerOfficialStreetName         GenericStringType235 `xml:"SellerOfficialStreetName"`
	SellerOfficialTownName           GenericStringType235 `xml:"SellerOfficialTownName"`
	SellerOfficialPostCodeIdentifier GenericStringType235 `xml:"SellerOfficialPostCodeIdentifier"`
	SellerOfficialCountrySubdivision GenericStringType235 `xml:"SellerOfficialCountrySubdivision,omitempty"`
	CountryCode                      CountryCodeType      `xml:"CountryCode,omitempty"`
	CountryName                      string               `xml:"CountryName,omitempty"`
}

type SellerPartyDetailsType struct {
	SellerPartyIdentifier            PartyLegalRegIdType            `xml:"SellerPartyIdentifier,omitempty"`
	SellerPartyIdentifierUrlText     string                         `xml:"SellerPartyIdentifierUrlText,omitempty"`
	SellerOrganisationName           []GenericStringType270         `xml:"SellerOrganisationName"`
	SellerOrganisationTradingName    GenericStringType270           `xml:"SellerOrganisationTradingName,omitempty"`
	SellerOrganisationDepartment     []string                       `xml:"SellerOrganisationDepartment,omitempty"`
	SellerOrganisationTaxCode        string                         `xml:"SellerOrganisationTaxCode,omitempty"`
	SellerOrganisationTaxCodeUrlText string                         `xml:"SellerOrganisationTaxCodeUrlText,omitempty"`
	SellerCode                       []PartyIdentifierType          `xml:"SellerCode,omitempty"`
	SellerPostalAddressDetails       SellerPostalAddressDetailsType `xml:"SellerPostalAddressDetails,omitempty"`
}

type SellerPostalAddressDetailsType struct {
	SellerStreetName              []GenericStringType235 `xml:"SellerStreetName"`
	SellerTownName                GenericStringType235   `xml:"SellerTownName"`
	SellerPostCodeIdentifier      GenericStringType235   `xml:"SellerPostCodeIdentifier"`
	SellerCountrySubdivision      GenericStringType235   `xml:"SellerCountrySubdivision,omitempty"`
	CountryCode                   CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                   string                 `xml:"CountryName,omitempty"`
	SellerPostOfficeBoxIdentifier string                 `xml:"SellerPostOfficeBoxIdentifier,omitempty"`
}

type ShipmentPartyDetails struct {
	ShipmentPartyIdentifier        PartyLegalRegIdType          `xml:"ShipmentPartyIdentifier,omitempty"`
	ShipmentOrganisationName       []GenericStringType235       `xml:"ShipmentOrganisationName"`
	ShipmentOrganisationDepartment []string                     `xml:"ShipmentOrganisationDepartment,omitempty"`
	ShipmentOrganisationTaxCode    string                       `xml:"ShipmentOrganisationTaxCode,omitempty"`
	ShipmentCode                   PartyIdentifierType          `xml:"ShipmentCode,omitempty"`
	ShipmentPostalAddressDetails   ShipmentPostalAddressDetails `xml:"ShipmentPostalAddressDetails,omitempty"`
	ShipmentSiteCode               string                       `xml:"ShipmentSiteCode,omitempty"`
}

type ShipmentPostalAddressDetails struct {
	ShipmentStreetName              []GenericStringType235 `xml:"ShipmentStreetName"`
	ShipmentTownName                GenericStringType235   `xml:"ShipmentTownName"`
	ShipmentPostCodeIdentifier      GenericStringType235   `xml:"ShipmentPostCodeIdentifier"`
	ShipmentCountrySubdivision      GenericStringType235   `xml:"ShipmentCountrySubdivision,omitempty"`
	CountryCode                     CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                     string                 `xml:"CountryName,omitempty"`
	ShipmentPostOfficeBoxIdentifier string                 `xml:"ShipmentPostOfficeBoxIdentifier,omitempty"`
}

type SpecificationDetailsType struct {
	SpecificationFreeText        []string                         `xml:"SpecificationFreeText,omitempty"`
	ExternalSpecificationDetails ExternalSpecificationDetailsType `xml:"ExternalSpecificationDetails,omitempty"`
}

type SubInvoiceRowType struct {
	SubIdentifier                             string                             `xml:"SubIdentifier,omitempty"`
	SubRowPositionIdentifier                  string                             `xml:"SubRowPositionIdentifier,omitempty"`
	SubInvoicedObjectID                       InvoicedObjectIDType               `xml:"SubInvoicedObjectID,omitempty"`
	SubArticleIdentifier                      string                             `xml:"SubArticleIdentifier,omitempty"`
	SubArticleGroupIdentifier                 []ArticleGroupIdentifierType       `xml:"SubArticleGroupIdentifier,omitempty"`
	SubArticleName                            string                             `xml:"SubArticleName,omitempty"`
	SubArticleDescription                     string                             `xml:"SubArticleDescription,omitempty"`
	SubArticleInfoUrlText                     string                             `xml:"SubArticleInfoUrlText,omitempty"`
	SubBuyerArticleIdentifier                 string                             `xml:"SubBuyerArticleIdentifier,omitempty"`
	SubEanCode                                EanCodeType                        `xml:"SubEanCode,omitempty"`
	SubRowRegistrationNumberIdentifier        string                             `xml:"SubRowRegistrationNumberIdentifier,omitempty"`
	SubSerialNumberIdentifier                 string                             `xml:"SubSerialNumberIdentifier,omitempty"`
	SubRowActionCode                          string                             `xml:"SubRowActionCode,omitempty"`
	SubRowDefinitionDetails                   []SubRowDefinitionDetails          `xml:"SubRowDefinitionDetails,omitempty"`
	SubOfferedQuantity                        []QuantityType014                  `xml:"SubOfferedQuantity,omitempty"`
	SubDeliveredQuantity                      []QuantityType014                  `xml:"SubDeliveredQuantity,omitempty"`
	SubOrderedQuantity                        QuantityType014                    `xml:"SubOrderedQuantity,omitempty"`
	SubConfirmedQuantity                      QuantityType014                    `xml:"SubConfirmedQuantity,omitempty"`
	SubPostDeliveredQuantity                  QuantityType014                    `xml:"SubPostDeliveredQuantity,omitempty"`
	SubInvoicedQuantity                       []QuantityType014                  `xml:"SubInvoicedQuantity,omitempty"`
	SubCreditRequestedQuantity                QuantityType014                    `xml:"SubCreditRequestedQuantity,omitempty"`
	SubReturnedQuantity                       QuantityType014                    `xml:"SubReturnedQuantity,omitempty"`
	SubStartDate                              Date                               `xml:"SubStartDate,omitempty"`
	SubEndDate                                Date                               `xml:"SubEndDate,omitempty"`
	SubUnitPriceAmount                        UnitAmount                         `xml:"SubUnitPriceAmount,omitempty"`
	SubUnitPriceDiscountAmount                UnitAmount                         `xml:"SubUnitPriceDiscountAmount,omitempty"`
	SubUnitPriceNetAmount                     UnitAmount                         `xml:"SubUnitPriceNetAmount,omitempty"`
	SubUnitPriceVatIncludedAmount             UnitAmount                         `xml:"SubUnitPriceVatIncludedAmount,omitempty"`
	SubUnitPriceBaseQuantity                  QuantityType014                    `xml:"SubUnitPriceBaseQuantity,omitempty"`
	SubRowIdentifier                          string                             `xml:"SubRowIdentifier,omitempty"`
	SubRowIdentifierUrlText                   string                             `xml:"SubRowIdentifierUrlText,omitempty"`
	SubRowIdentifierDate                      Date                               `xml:"SubRowIdentifierDate,omitempty"`
	SubRowOrdererName                         string                             `xml:"SubRowOrdererName,omitempty"`
	SubRowSalesPersonName                     string                             `xml:"SubRowSalesPersonName,omitempty"`
	SubRowOrderConfirmationIdentifier         string                             `xml:"SubRowOrderConfirmationIdentifier,omitempty"`
	SubRowOrderConfirmationDate               Date                               `xml:"SubRowOrderConfirmationDate,omitempty"`
	SubOriginalInvoiceNumber                  GenericStringType120               `xml:"SubOriginalInvoiceNumber,omitempty"`
	SubOriginalInvoiceDate                    Date                               `xml:"SubOriginalInvoiceDate,omitempty"`
	SubOriginalInvoiceReference               []OriginalInvoiceReferenceType     `xml:"SubOriginalInvoiceReference,omitempty"`
	SubRowDeliveryIdentifier                  string                             `xml:"SubRowDeliveryIdentifier,omitempty"`
	SubRowDeliveryIdentifierUrlText           string                             `xml:"SubRowDeliveryIdentifierUrlText,omitempty"`
	SubRowDeliveryDate                        Date                               `xml:"SubRowDeliveryDate,omitempty"`
	SubRowQuotationIdentifier                 string                             `xml:"SubRowQuotationIdentifier,omitempty"`
	SubRowQuotationIdentifierUrlText          string                             `xml:"SubRowQuotationIdentifierUrlText,omitempty"`
	SubRowAgreementIdentifier                 string                             `xml:"SubRowAgreementIdentifier,omitempty"`
	SubRowAgreementIdentifierUrlText          string                             `xml:"SubRowAgreementIdentifierUrlText,omitempty"`
	SubRowRequestOfQuotationIdentifier        string                             `xml:"SubRowRequestOfQuotationIdentifier,omitempty"`
	SubRowRequestOfQuotationIdentifierUrlText string                             `xml:"SubRowRequestOfQuotationIdentifierUrlText,omitempty"`
	SubRowPriceListIdentifier                 string                             `xml:"SubRowPriceListIdentifier,omitempty"`
	SubRowPriceListIdentifierUrlText          string                             `xml:"SubRowPriceListIdentifierUrlText,omitempty"`
	SubRowBuyerReferenceIdentifier            string                             `xml:"SubRowBuyerReferenceIdentifier,omitempty"`
	SubRowProjectReferenceIdentifier          string                             `xml:"SubRowProjectReferenceIdentifier,omitempty"`
	SubRowOverDuePaymentDetails               SubRowOverDuePaymentDetails        `xml:"SubRowOverDuePaymentDetails,omitempty"`
	SubRowAnyPartyDetails                     []SubRowAnyPartyDetails            `xml:"SubRowAnyPartyDetails,omitempty"`
	SubRowDeliveryDetails                     SubRowDeliveryDetailsType          `xml:"SubRowDeliveryDetails,omitempty"`
	SubRowShortProposedAccountIdentifier      string                             `xml:"SubRowShortProposedAccountIdentifier,omitempty"`
	SubRowNormalProposedAccountIdentifier     string                             `xml:"SubRowNormalProposedAccountIdentifier,omitempty"`
	SubRowProposedAccountText                 string                             `xml:"SubRowProposedAccountText,omitempty"`
	SubRowAccountDimensionText                string                             `xml:"SubRowAccountDimensionText,omitempty"`
	SubRowSellerAccountText                   string                             `xml:"SubRowSellerAccountText,omitempty"`
	SubRowFreeText                            []string                           `xml:"SubRowFreeText,omitempty"`
	SubRowUsedQuantity                        QuantityType014                    `xml:"SubRowUsedQuantity,omitempty"`
	SubRowPreviousMeterReadingDate            Date                               `xml:"SubRowPreviousMeterReadingDate,omitempty"`
	SubRowLatestMeterReadingDate              Date                               `xml:"SubRowLatestMeterReadingDate,omitempty"`
	SubRowCalculatedQuantity                  QuantityType014                    `xml:"SubRowCalculatedQuantity,omitempty"`
	SubRowAveragePriceAmount                  Amount                             `xml:"SubRowAveragePriceAmount,omitempty"`
	SubRowDiscountPercent                     Percentage                         `xml:"SubRowDiscountPercent,omitempty"`
	SubRowDiscountAmount                      Amount                             `xml:"SubRowDiscountAmount,omitempty"`
	SubRowDiscountBaseAmount                  Amount                             `xml:"SubRowDiscountBaseAmount,omitempty"`
	SubRowDiscountTypeCode                    string                             `xml:"SubRowDiscountTypeCode,omitempty"`
	SubRowDiscountTypeText                    string                             `xml:"SubRowDiscountTypeText,omitempty"`
	SubRowProgressiveDiscountDetails          []SubRowProgressiveDiscountDetails `xml:"SubRowProgressiveDiscountDetails,omitempty"`
	SubRowChargeDetails                       []RowChargeDetailsType             `xml:"SubRowChargeDetails,omitempty"`
	SubRowVatRatePercent                      Percentage                         `xml:"SubRowVatRatePercent,omitempty"`
	SubRowVatCode                             string                             `xml:"SubRowVatCode,omitempty"`
	SubRowVatAmount                           Amount                             `xml:"SubRowVatAmount,omitempty"`
	SubRowVatExcludedAmount                   Amount                             `xml:"SubRowVatExcludedAmount,omitempty"`
	SubRowAmount                              Amount                             `xml:"SubRowAmount,omitempty"`
	SubRowTransactionDetails                  TransactionDetailsType             `xml:"SubRowTransactionDetails,omitempty"`
}

type SubRowAnyPartyDetails struct {
	SubRowAnyPartyText                   Anypartytexttype035                `xml:"SubRowAnyPartyText"`
	SubRowAnyPartyIdentifier             PartyLegalRegIdType                `xml:"SubRowAnyPartyIdentifier,omitempty"`
	SubRowAnyPartyOrganisationName       []GenericStringType235             `xml:"SubRowAnyPartyOrganisationName"`
	SubRowAnyPartyOrganisationDepartment []string                           `xml:"SubRowAnyPartyOrganisationDepartment,omitempty"`
	SubRowAnyPartyOrganisationTaxCode    string                             `xml:"SubRowAnyPartyOrganisationTaxCode,omitempty"`
	SubRowAnyPartyCode                   PartyIdentifierType                `xml:"SubRowAnyPartyCode,omitempty"`
	SubRowAnyPartyPostalAddressDetails   SubRowAnyPartyPostalAddressDetails `xml:"SubRowAnyPartyPostalAddressDetails,omitempty"`
	SubRowAnyPartyOrganisationUnitNumber string                             `xml:"SubRowAnyPartyOrganisationUnitNumber,omitempty"`
	SubRowAnyPartySiteCode               string                             `xml:"SubRowAnyPartySiteCode,omitempty"`
}

type SubRowAnyPartyPostalAddressDetails struct {
	SubRowAnyPartyStreetName              []GenericStringType235 `xml:"SubRowAnyPartyStreetName"`
	SubRowAnyPartyTownName                GenericStringType235   `xml:"SubRowAnyPartyTownName"`
	SubRowAnyPartyPostCodeIdentifier      GenericStringType235   `xml:"SubRowAnyPartyPostCodeIdentifier"`
	SubRowAnyPartyCountrySubdivision      GenericStringType235   `xml:"SubRowAnyPartyCountrySubdivision,omitempty"`
	CountryCode                           CountryCodeType        `xml:"CountryCode,omitempty"`
	CountryName                           string                 `xml:"CountryName,omitempty"`
	SubRowAnyPartyPostOfficeBoxIdentifier string                 `xml:"SubRowAnyPartyPostOfficeBoxIdentifier,omitempty"`
}

type SubRowDefinitionDetails struct {
	SubRowDefinitionHeaderText SubRowDefinitionHeaderText `xml:"SubRowDefinitionHeaderText"`
	SubRowDefinitionValue      QuantityType070            `xml:"SubRowDefinitionValue,omitempty"`
}

type SubRowDefinitionHeaderText struct {
	Value          string              `xml:",chardata"`
	DefinitionCode GenericTokenType120 `xml:"DefinitionCode,attr,omitempty"`
}

type SubRowDeliveryDetailsType struct {
	SubRowTerminalAddressText            string               `xml:"SubRowTerminalAddressText,omitempty"`
	SubRowWaybillIdentifier              string               `xml:"SubRowWaybillIdentifier,omitempty"`
	SubRowWaybillTypeCode                string               `xml:"SubRowWaybillTypeCode,omitempty"`
	SubRowClearanceIdentifier            string               `xml:"SubRowClearanceIdentifier,omitempty"`
	SubRowDeliveryNoteIdentifier         string               `xml:"SubRowDeliveryNoteIdentifier,omitempty"`
	SubRowDelivererIdentifier            string               `xml:"SubRowDelivererIdentifier,omitempty"`
	SubRowDelivererName                  []string             `xml:"SubRowDelivererName,omitempty"`
	SubRowDelivererCountrySubdivision    GenericStringType235 `xml:"SubRowDelivererCountrySubdivision,omitempty"`
	SubRowDelivererCountryCode           CountryCodeType      `xml:"SubRowDelivererCountryCode,omitempty"`
	SubRowDelivererCountryName           string               `xml:"SubRowDelivererCountryName,omitempty"`
	SubRowPlaceOfDischarge               string               `xml:"SubRowPlaceOfDischarge,omitempty"`
	SubRowFinalDestinationName           []string             `xml:"SubRowFinalDestinationName,omitempty"`
	SubRowCustomsInfo                    CustomsInfoType      `xml:"SubRowCustomsInfo,omitempty"`
	SubRowManufacturerArticleIdentifier  string               `xml:"SubRowManufacturerArticleIdentifier,omitempty"`
	SubRowManufacturerIdentifier         string               `xml:"SubRowManufacturerIdentifier,omitempty"`
	SubRowManufacturerName               []string             `xml:"SubRowManufacturerName,omitempty"`
	SubRowManufacturerCountrySubdivision GenericStringType235 `xml:"SubRowManufacturerCountrySubdivision,omitempty"`
	SubRowManufacturerCountryCode        CountryCodeType      `xml:"SubRowManufacturerCountryCode,omitempty"`
	SubRowManufacturerCountryName        string               `xml:"SubRowManufacturerCountryName,omitempty"`
	SubRowManufacturerOrderIdentifier    string               `xml:"SubRowManufacturerOrderIdentifier,omitempty"`
	SubRowPackageDetails                 SubRowPackageDetails `xml:"SubRowPackageDetails,omitempty"`
}

type SubRowOverDuePaymentDetails struct {
	SubRowOriginalInvoiceIdentifier           string          `xml:"SubRowOriginalInvoiceIdentifier,omitempty"`
	SubRowOriginalInvoiceDate                 Date            `xml:"SubRowOriginalInvoiceDate,omitempty"`
	SubRowOriginalDueDate                     Date            `xml:"SubRowOriginalDueDate,omitempty"`
	SubRowOriginalInvoiceTotalAmount          Amount          `xml:"SubRowOriginalInvoiceTotalAmount,omitempty"`
	SubRowOriginalEpiRemittanceInfoIdentifier string          `xml:"SubRowOriginalEpiRemittanceInfoIdentifier,omitempty"`
	SubRowPaidVatExcludedAmount               Amount          `xml:"SubRowPaidVatExcludedAmount,omitempty"`
	SubRowPaidVatIncludedAmount               Amount          `xml:"SubRowPaidVatIncludedAmount,omitempty"`
	SubRowPaidDate                            Date            `xml:"SubRowPaidDate,omitempty"`
	SubRowUnPaidVatExcludedAmount             Amount          `xml:"SubRowUnPaidVatExcludedAmount,omitempty"`
	SubRowUnPaidVatIncludedAmount             Amount          `xml:"SubRowUnPaidVatIncludedAmount,omitempty"`
	SubRowCollectionDate                      Date            `xml:"SubRowCollectionDate,omitempty"`
	SubRowCollectionQuantity                  QuantityType014 `xml:"SubRowCollectionQuantity,omitempty"`
	SubRowCollectionChargeAmount              Amount          `xml:"SubRowCollectionChargeAmount,omitempty"`
	SubRowInterestRate                        Percentage      `xml:"SubRowInterestRate,omitempty"`
	SubRowInterestStartDate                   Date            `xml:"SubRowInterestStartDate,omitempty"`
	SubRowInterestEndDate                     Date            `xml:"SubRowInterestEndDate,omitempty"`
	SubRowInterestPeriodText                  string          `xml:"SubRowInterestPeriodText,omitempty"`
	SubRowInterestDateNumber                  string          `xml:"SubRowInterestDateNumber,omitempty"`
	SubRowInterestChargeAmount                Amount          `xml:"SubRowInterestChargeAmount,omitempty"`
	SubRowInterestChargeVatAmount             Amount          `xml:"SubRowInterestChargeVatAmount,omitempty"`
}

type SubRowPackageDetails struct {
	SubRowPackageLength             QuantityType014 `xml:"SubRowPackageLength,omitempty"`
	SubRowPackageWidth              QuantityType014 `xml:"SubRowPackageWidth,omitempty"`
	SubRowPackageHeight             QuantityType014 `xml:"SubRowPackageHeight,omitempty"`
	SubRowPackageWeight             QuantityType014 `xml:"SubRowPackageWeight,omitempty"`
	SubRowPackageNetWeight          QuantityType014 `xml:"SubRowPackageNetWeight,omitempty"`
	SubRowPackageVolume             QuantityType014 `xml:"SubRowPackageVolume,omitempty"`
	SubRowTransportCarriageQuantity QuantityType014 `xml:"SubRowTransportCarriageQuantity,omitempty"`
}

type SubRowProgressiveDiscountDetails struct {
	SubRowDiscountPercent    Percentage `xml:"SubRowDiscountPercent,omitempty"`
	SubRowDiscountAmount     Amount     `xml:"SubRowDiscountAmount,omitempty"`
	SubRowDiscountBaseAmount Amount     `xml:"SubRowDiscountBaseAmount,omitempty"`
	SubRowDiscountTypeCode   string     `xml:"SubRowDiscountTypeCode,omitempty"`
	SubRowDiscountTypeText   string     `xml:"SubRowDiscountTypeText,omitempty"`
}

type TransactionDetailsType struct {
	OtherCurrencyAmount Amount       `xml:"OtherCurrencyAmount,omitempty"`
	ExchangeRate        ExchangeRate `xml:"ExchangeRate,omitempty"`
	ExchangeDate        Date         `xml:"ExchangeDate,omitempty"`
}

type UnitAmount struct {
	UnitAmountType           UnitAmountType           `xml:",chardata"`
	AmountCurrencyIdentifier AmountCurrencyIdentifier `xml:"AmountCurrencyIdentifier,attr"`
	UnitPriceUnitCode        string                   `xml:"UnitPriceUnitCode,attr,omitempty"`
}

// Must match the pattern -?[0-9]{1,15}(,[0-9]{2,5})?
type UnitAmountType string

type UnitAmountUN struct {
	UnitAmountType           UnitAmountType           `xml:",chardata"`
	QuantityUnitCodeUN       UnitCodeUN               `xml:"QuantityUnitCodeUN,attr,omitempty"`
	AmountCurrencyIdentifier AmountCurrencyIdentifier `xml:"AmountCurrencyIdentifier,attr"`
	UnitPriceUnitCode        string                   `xml:"UnitPriceUnitCode,attr,omitempty"`
}

// Must match the pattern [A-Z0-9]{2,3}
type UnitCodeUN string

// Must be at least 1 items long
type Untdid13AZ09 string

type Untdid5305 string

// Must match the pattern [A-Za-z0-9\-]{1,20}
type VatExReasonCodeType string

type VatPointType struct {
	VatPointDate     Date   `xml:"VatPointDate"`
	VatPointDateCode string `xml:"VatPointDateCode"`
}

type VatSpecificationDetailsType struct {
	VatBaseAmount          Amount              `xml:"VatBaseAmount,omitempty"`
	VatRatePercent         Percentage          `xml:"VatRatePercent,omitempty"`
	VatCode                string              `xml:"VatCode,omitempty"`
	VatRateAmount          Amount              `xml:"VatRateAmount,omitempty"`
	VatFreeText            []string            `xml:"VatFreeText,omitempty"`
	VatExemptionReasonCode VatExReasonCodeType `xml:"VatExemptionReasonCode,omitempty"`
}

// May be one of 1.3, 2.0, 2.01, 3.0
type Version string

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}
