package xsd

var Finvoice30xsd = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!-- Modified 30.11.2020 -->
<!-- For Finvoice version 3.0 -->
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema" elementFormDefault="qualified" attributeFormDefault="unqualified">
  <xs:element name="Finvoice">
    <xs:complexType>
      <xs:sequence>
        <xs:element name="MessageTransmissionDetails" type="MessageTransmissionDetailsType" minOccurs="0"/>
        <xs:element name="SellerPartyDetails" type="SellerPartyDetailsType"/>
        <xs:element name="SellerOrganisationUnitNumber" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="SellerSiteCode" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="SellerContactPersonName" type="genericStringType0_70" minOccurs="0"/>
        <xs:element name="SellerContactPersonFunction" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="SellerContactPersonDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="SellerCommunicationDetails" type="SellerCommunicationDetailsType" minOccurs="0"/>
        <xs:element name="SellerInformationDetails" type="SellerInformationDetailsType" minOccurs="0"/>
        <xs:element name="InvoiceSenderPartyDetails" type="InvoiceSenderPartyDetailsType" minOccurs="0"/>
        <xs:element name="InvoiceRecipientPartyDetails" type="InvoiceRecipientPartyDetailsType" minOccurs="0"/>
        <xs:element name="InvoiceRecipientOrganisationUnitNumber" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="InvoiceRecipientSiteCode" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="InvoiceRecipientContactPersonName" type="genericStringType0_70" minOccurs="0"/>
        <xs:element name="InvoiceRecipientContactPersonFunction" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="InvoiceRecipientContactPersonDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="InvoiceRecipientLanguageCode" type="LanguageCodeType" minOccurs="0"/>
        <xs:element name="InvoiceRecipientCommunicationDetails" type="InvoiceRecipientCommunicationDetailsType" minOccurs="0"/>
        <xs:element name="BuyerPartyDetails" type="BuyerPartyDetailsType"/>
        <xs:element name="BuyerOrganisationUnitNumber" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="BuyerSiteCode" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="BuyerContactPersonName" type="genericStringType0_70" minOccurs="0"/>
        <xs:element name="BuyerContactPersonFunction" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="BuyerContactPersonDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="BuyerCommunicationDetails" type="BuyerCommunicationDetailsType" minOccurs="0"/>
        <xs:element name="DeliveryPartyDetails" type="DeliveryPartyDetailsType" minOccurs="0"/>
        <xs:element name="DeliveryOrganisationUnitNumber" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="DeliverySiteCode" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="DeliveryContactPersonName" type="genericStringType0_70" minOccurs="0"/>
        <xs:element name="DeliveryContactPersonFunction" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="DeliveryContactPersonDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
        <xs:element name="DeliveryCommunicationDetails" type="DeliveryCommunicationDetailsType" minOccurs="0"/>
        <xs:element name="DeliveryDetails" type="DeliveryDetailsType" minOccurs="0"/>
        <xs:element name="AnyPartyDetails" type="AnyPartyDetailsType" minOccurs="0" maxOccurs="unbounded"/>
        <xs:element name="InvoiceDetails" type="InvoiceDetailsType"/>
        <xs:element name="PaymentCardInfo" type="PaymentCardInfoType" minOccurs="0"/>
        <xs:element name="DirectDebitInfo" type="DirectDebitInfoType" minOccurs="0"/>
        <xs:element name="PaymentStatusDetails" type="PaymentStatusDetailsType" minOccurs="0"/>
        <xs:element name="PartialPaymentDetails" type="PartialPaymentDetailsType" minOccurs="0" maxOccurs="unbounded"/>
        <xs:element name="FactoringAgreementDetails" type="FactoringAgreementDetailsType" minOccurs="0"/>
        <xs:element name="VirtualBankBarcode" type="genericStringType0_512" minOccurs="0"/>
        <xs:element name="InvoiceRow" type="InvoiceRowType" maxOccurs="unbounded"/>
        <xs:element name="SpecificationDetails" type="SpecificationDetailsType" minOccurs="0"/>
        <xs:element name="EpiDetails" type="EpiDetailsType"/>
        <xs:element name="InvoiceUrlNameText" type="genericStringType0_512" minOccurs="0" maxOccurs="unbounded"/>
        <xs:element name="InvoiceUrlText" type="genericStringType0_512" minOccurs="0" maxOccurs="unbounded"/>
        <xs:element name="StorageUrlText" type="genericStringType0_512" minOccurs="0"/>
        <xs:element name="LayOutIdentifier" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="InvoiceSegmentIdentifier" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="ControlChecksum" type="genericStringType0_512" minOccurs="0"/>
        <xs:element name="MessageChecksum" type="genericStringType0_512" minOccurs="0"/>
        <xs:element name="ControlStampText" type="genericStringType0_512" minOccurs="0"/>
        <xs:element name="AcceptanceStampText" type="genericStringType0_512" minOccurs="0"/>
        <xs:element name="OriginalInvoiceFormat" type="genericStringType0_35" minOccurs="0"/>
        <xs:element name="AttachmentMessageDetails" type="AttachmentMessageDetailsType" minOccurs="0"/>
      </xs:sequence>
      <xs:attribute name="Version" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:NMTOKEN">
            <xs:enumeration value="1.3"/>
            <xs:enumeration value="2.0"/>
            <xs:enumeration value="2.01"/>
            <xs:enumeration value="3.0"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
    </xs:complexType>
  </xs:element>
  <xs:complexType name="MessageTransmissionDetailsType">
    <xs:sequence>
      <xs:element name="MessageSenderDetails">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="FromIdentifier" type="ElectronicAddrIdType"/>
            <xs:element name="FromIntermediator" type="genericStringType2_35"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="MessageReceiverDetails">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="ToIdentifier" type="ElectronicAddrIdType"/>
            <xs:element name="ToIntermediator" type="genericStringType2_35"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="MessageDetails">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="MessageIdentifier" type="genericStringType2_48"/>
            <xs:element name="MessageTimeStamp" type="xs:dateTime"/>
            <xs:element name="RefToMessageIdentifier" type="genericStringType0_48" minOccurs="0"/>
            <xs:element name="ImplementationCode" type="genericStringType0_4" minOccurs="0"/>
            <xs:element name="SpecificationIdentifier" type="genericStringType1_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="AnyPartyDetailsType">
    <xs:sequence>
      <xs:element name="AnyPartyText" type="anypartytexttype0_35"/>
      <xs:element name="AnyPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="AnyPartyOrganisationName" type="genericStringType2_35" maxOccurs="2"/>
      <xs:element name="AnyPartyOrganisationDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="AnyPartyOrganisationTaxCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="AnyPartyCode" type="PartyIdentifierType" minOccurs="0"/>
      <xs:element name="AnyPartyContactPersonName" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="AnyPartyContactPersonFunction" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="AnyPartyContactPersonDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="AnyPartyCommunicationDetails" type="AnyPartyCommunicationDetailsType" minOccurs="0"/>
      <xs:element name="AnyPartyPostalAddressDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="AnyPartyStreetName" type="genericStringType2_35" maxOccurs="3"/>
            <xs:element name="AnyPartyTownName" type="genericStringType2_35"/>
            <xs:element name="AnyPartyPostCodeIdentifier" type="genericStringType2_35"/>
            <xs:element name="AnyPartyCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
            <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
            <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="AnyPartyPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="AnyPartyOrganisationUnitNumber" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="AnyPartySiteCode" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="FactoringAgreementDetailsType">
    <xs:sequence>
      <xs:element name="FactoringAgreementIdentifier" type="genericStringType0_35"/>
      <xs:element name="TransmissionListIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="EndorsementClauseCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="FactoringTypeCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="FactoringFreeText" type="genericStringType0_70" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="FactoringPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="FactoringPartyName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="FactoringPartyPostalAddressDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="FactoringPartyStreetName" type="genericStringType2_35" maxOccurs="3"/>
            <xs:element name="FactoringPartyTownName" type="genericStringType2_35"/>
            <xs:element name="FactoringPartyPostCodeIdentifier" type="genericStringType2_35"/>
            <xs:element name="FactoringPartyCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
            <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
            <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="FactoringPartyPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="BuyerCommunicationDetailsType">
    <xs:sequence>
      <xs:element name="BuyerPhoneNumberIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="BuyerEmailaddressIdentifier" type="genericStringType0_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="BuyerPartyDetailsType">
    <xs:sequence>
      <xs:element name="BuyerPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="BuyerOrganisationName" type="genericStringType2_70" maxOccurs="unbounded"/>
      <xs:element name="BuyerOrganisationTradingName" type="genericStringType2_70" minOccurs="0"/>
      <xs:element name="BuyerOrganisationDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="BuyerOrganisationTaxCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="BuyerCode" type="PartyIdentifierType" minOccurs="0"/>
      <xs:element name="BuyerPostalAddressDetails" type="BuyerPostalAddressDetailsType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="BuyerPostalAddressDetailsType">
    <xs:sequence>
      <xs:element name="BuyerStreetName" type="genericStringType2_35" maxOccurs="3"/>
      <xs:element name="BuyerTownName" type="genericStringType2_35"/>
      <xs:element name="BuyerPostCodeIdentifier" type="genericStringType2_35"/>
      <xs:element name="BuyerCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="BuyerPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="DeliveryCommunicationDetailsType">
    <xs:sequence>
      <xs:element name="DeliveryPhoneNumberIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="DeliveryEmailaddressIdentifier" type="genericStringType0_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="DeliveryDetailsType">
    <xs:sequence>
      <xs:element name="DeliveryDate" type="date" minOccurs="0"/>
      <xs:element name="DeliveryPeriodDetails" type="DeliveryPeriodDetailsType" minOccurs="0"/>
      <xs:element name="ShipmentPartyDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="ShipmentPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
            <xs:element name="ShipmentOrganisationName" type="genericStringType2_35" maxOccurs="unbounded"/>
            <xs:element name="ShipmentOrganisationDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
            <xs:element name="ShipmentOrganisationTaxCode" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="ShipmentCode" type="PartyIdentifierType" minOccurs="0"/>
            <xs:element name="ShipmentPostalAddressDetails" minOccurs="0">
              <xs:complexType>
                <xs:sequence>
                  <xs:element name="ShipmentStreetName" type="genericStringType2_35" maxOccurs="3"/>
                  <xs:element name="ShipmentTownName" type="genericStringType2_35"/>
                  <xs:element name="ShipmentPostCodeIdentifier" type="genericStringType2_35"/>
                  <xs:element name="ShipmentCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
                  <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
                  <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
                  <xs:element name="ShipmentPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
                </xs:sequence>
              </xs:complexType>
            </xs:element>
            <xs:element name="ShipmentSiteCode" type="genericStringType0_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="DeliveryMethodText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="DeliveryTermsText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="DeliveryTermsCode" type="genericStringType1_4" minOccurs="0"/>
      <xs:element name="TerminalAddressText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="WaybillIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="WaybillTypeCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="ClearanceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="DeliveryNoteIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="DelivererIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="DelivererName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="DelivererCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="DelivererCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="DelivererCountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="ModeOfTransportIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="CarrierName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="VesselName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="LocationIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="TransportInformationDate" type="date" minOccurs="0"/>
      <xs:element name="CountryOfOrigin" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="CountryOfDestinationName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="DestinationCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="PlaceOfDischarge" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="FinalDestinationName" type="DestinationNameType" minOccurs="0" maxOccurs="3"/>
      <xs:element name="ManufacturerIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="ManufacturerName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="ManufacturerCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="ManufacturerCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="ManufacturerCountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="ManufacturerOrderIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="PackageDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="PackageLength" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="PackageWidth" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="PackageHeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="PackageWeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="PackageNetWeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="PackageVolume" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="TransportCarriageQuantity" type="QuantityType0_14" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="DeliveryPartyDetailsType">
    <xs:sequence>
      <xs:element name="DeliveryPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="DeliveryOrganisationName" type="genericStringType2_35" maxOccurs="unbounded"/>
      <xs:element name="DeliveryOrganisationDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="DeliveryOrganisationTaxCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="DeliveryCode" type="PartyIdentifierType" minOccurs="0"/>
      <xs:element name="DeliveryPostalAddressDetails" type="DeliveryPostalAddressDetailsType"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="DeliveryPeriodDetailsType">
    <xs:sequence>
      <xs:element name="StartDate" type="date"/>
      <xs:element name="EndDate" type="date"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="DeliveryPostalAddressDetailsType">
    <xs:sequence>
      <xs:element name="DeliveryStreetName" type="genericStringType2_35" maxOccurs="3"/>
      <xs:element name="DeliveryTownName" type="genericStringType2_35"/>
      <xs:element name="DeliveryPostCodeIdentifier" type="genericStringType2_35"/>
      <xs:element name="DeliveryCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="DeliveryPostofficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="EpiAccountIDType">
    <xs:simpleContent>
      <xs:extension base="genericNMtokenType1_34">
        <xs:attribute name="IdentificationSchemeName" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="IBAN"/>
              <xs:enumeration value="BBAN"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="EpiBfiIdentifierType">
    <xs:simpleContent>
      <xs:extension base="genericNMtokenType8_11">
        <xs:attribute name="IdentificationSchemeName" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="BIC"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="EpiDetailsType">
    <xs:sequence>
      <xs:element name="EpiIdentificationDetails" type="EpiIdentificationDetailsType"/>
      <xs:element name="EpiPartyDetails" type="EpiPartyDetailsType"/>
      <xs:element name="EpiPaymentInstructionDetails" type="EpiPaymentInstructionDetailsType"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="EpiIdentificationDetailsType">
    <xs:sequence>
      <xs:element name="EpiDate" type="date"/>
      <xs:element name="EpiReference" type="genericNMtokenType0_35"/>
      <xs:element name="EpiUrl" type="genericNMtokenType0_512" minOccurs="0"/>
      <xs:element name="EpiEmail" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="EpiOrderInfo" type="genericTokenType0_70" minOccurs="0" maxOccurs="7"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="EpiPartyDetailsType">
    <xs:sequence>
      <xs:element name="EpiBfiPartyDetails" type="EpiBfiPartyDetailsType"/>
      <xs:element name="EpiBeneficiaryPartyDetails" type="EpiBeneficiaryPartyDetailsType"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="EpiBfiPartyDetailsType">
    <xs:sequence>
      <xs:element name="EpiBfiIdentifier" type="EpiBfiIdentifierType" minOccurs="0"/>
      <xs:element name="EpiBfiName" type="genericStringType1_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="EpiBeneficiaryPartyDetailsType">
    <xs:sequence>
      <xs:element name="EpiNameAddressDetails" type="genericTokenType2_35" minOccurs="0"/>
      <xs:element name="EpiBei" type="genericNMtokenType8_11" minOccurs="0"/>
      <xs:element name="EpiAccountID" type="EpiAccountIDType"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="EpiPaymentInstructionDetailsType">
    <xs:sequence>
      <xs:element name="EpiPaymentInstructionId" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="EpiTransactionTypeCode" type="genericTokenType3" minOccurs="0"/>
      <xs:element name="EpiInstructionCode" type="genericNMtokenType0_35" minOccurs="0"/>
      <xs:element name="EpiRemittanceInfoIdentifier" type="EpiRemittanceInfoIdentifierType" minOccurs="0"/>
      <xs:element name="EpiInstructedAmount" type="epiAmount"/>
      <xs:element name="EpiCharge" type="EpiChargeType"/>
      <xs:element name="EpiDateOptionDate" type="date"/>
      <xs:element name="EpiPaymentMeansCode" type="untdid4461" minOccurs="0"/>
      <xs:element name="EpiPaymentMeansText" type="genericStringType1_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="EpiRemittanceInfoIdentifierType">
    <xs:simpleContent>
      <xs:extension base="EpiRemittanceInfoIdentifierPattern">
        <xs:attribute name="IdentificationSchemeName">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="SPY"/>
              <xs:enumeration value="ISO"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="EpiRemittanceInfoIdentifierPattern">
    <xs:restriction base="xs:NMTOKEN">
      <xs:pattern value="([0-9]{2,20})|(RF[0-9][0-9][0-9A-Za-z]{1,21})"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="EpiChargeType">
    <xs:simpleContent>
      <xs:extension base="xs:token">
        <xs:attribute name="ChargeOption" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="SHA"/>
              <xs:enumeration value="OUR"/>
              <xs:enumeration value="BEN"/>
              <xs:enumeration value="SLEV"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="InvoiceDetailsType">
    <xs:sequence>
      <xs:element name="InvoiceTypeCode" type="InvoiceTypeCodeTypeFI"/>
      <xs:element name="InvoiceTypeCodeUN" type="untdid1001" minOccurs="0"/>
      <xs:element name="InvoiceTypeText" type="genericStringType1_35"/>
      <xs:element name="InvoiceClassification" type="InvoiceClassificationType" minOccurs="0"/>
      <xs:element name="OriginCode" type="OriginCodeType"/>
      <xs:element name="OriginText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="InvoicedObjectID" type="InvoicedObjectIDType" minOccurs="0"/>
      <xs:element name="InvoiceNumber" type="genericStringType1_20"/>
      <xs:element name="InvoiceDate" type="date"/>
      <xs:element name="OriginalInvoiceNumber" type="genericStringType1_20" minOccurs="0"/>
      <xs:element name="OriginalInvoiceDate" type="date" minOccurs="0"/>
      <xs:element name="OriginalInvoiceReference" type="OriginalInvoiceReferenceType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="InvoicingPeriodStartDate" type="date" minOccurs="0"/>
      <xs:element name="InvoicingPeriodEndDate" type="date" minOccurs="0"/>
      <xs:element name="SellerReferenceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SellerReferenceIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="BuyersSellerIdentifier" type="PartyIdentifierType" minOccurs="0"/>
      <xs:element name="SellersBuyerIdentifier" type="PartyIdentifierType" minOccurs="0"/>
      <xs:element name="OrderIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="OrderIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="OrderDate" type="date" minOccurs="0"/>
      <xs:element name="OrdererName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SalesPersonName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="OrderConfirmationIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="OrderConfirmationDate" type="date" minOccurs="0"/>
      <xs:element name="AgreementIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="AgreementIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="AgreementTypeText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="AgreementTypeCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="AgreementDate" type="date" minOccurs="0"/>
      <xs:element name="NotificationIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="NotificationDate" type="date" minOccurs="0"/>
      <xs:element name="RegistrationNumberIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="ControllerIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="ControllerName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="ControlDate" type="date" minOccurs="0"/>
      <xs:element name="BuyerReferenceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="ProjectReferenceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="DefinitionDetails" minOccurs="0" maxOccurs="unbounded">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="DefinitionHeaderText">
              <xs:complexType>
                <xs:simpleContent>
                  <xs:extension base="genericStringType0_70">
                    <xs:attribute name="DefinitionCode" type="genericTokenType1_20"/>
                  </xs:extension>
                </xs:simpleContent>
              </xs:complexType>
            </xs:element>
            <xs:element name="DefinitionValue" type="QuantityType0_70" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="RowsTotalVatExcludedAmount" minOccurs="0" type="amount"/>
      <xs:element name="DiscountsTotalVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="ChargesTotalVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="InvoiceTotalVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="InvoiceTotalVatAmount" type="amount" minOccurs="0"/>
      <xs:element name="InvoiceTotalVatAccountingAmount" type="amount" minOccurs="0"/>
      <xs:element name="InvoiceTotalVatIncludedAmount" type="amount"/>
      <xs:element name="InvoiceTotalRoundoffAmount" type="amount" minOccurs="0"/>
      <xs:element name="InvoicePaidAmount" type="amount" minOccurs="0"/>
      <xs:element name="ExchangeRate" type="exchangeRate" minOccurs="0"/>
      <xs:element name="OtherCurrencyAmountVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="OtherCurrencyAmountVatIncludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="CreditLimitAmount" type="amount" minOccurs="0"/>
      <xs:element name="CreditInterestPercent" type="percentage" minOccurs="0"/>
      <xs:element name="OperationLimitAmount" type="amount" minOccurs="0"/>
      <xs:element name="MonthlyAmount" type="amount" minOccurs="0"/>
      <xs:element name="ShortProposedAccountIdentifier" type="genericNMtokenType0_4" minOccurs="0"/>
      <xs:element name="NormalProposedAccountIdentifier" type="genericNMtokenType0_4" minOccurs="0"/>
      <xs:element name="ProposedAccountText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="AccountDimensionText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerAccountText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="VatPoint" type="VatPointType" minOccurs="0"/>
      <xs:element name="VatSpecificationDetails" type="VatSpecificationDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="InvoiceFreeText" type="genericStringType0_512" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="InvoiceVatFreeText" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="PaymentTermsDetails" type="PaymentTermsDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="DiscountDetails" type="DiscountDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="ChargeDetails" type="InvoiceChargeDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="TenderReference" type="genericStringType1_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="OriginalInvoiceReferenceType">
    <xs:sequence>
      <xs:element name="InvoiceNumber" type="genericStringType1_20" minOccurs="0"/>
      <xs:element name="InvoiceDate" type="date" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="PaymentCardInfoType">
    <xs:sequence>
      <xs:element name="PrimaryAccountNumber" type="genericStringType1_19"/>
      <xs:element name="CardHolderName" type="genericStringType1_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="DirectDebitInfoType">
    <xs:sequence>
      <xs:element name="MandateReference" type="genericStringType1_35" minOccurs="0"/>
      <xs:element name="CreditorIdentifier" type="genericStringType1_35" minOccurs="0"/>
      <xs:element name="DebitedAccountID" type="EpiAccountIDType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="InvoiceRecipientCommunicationDetailsType">
    <xs:sequence>
      <xs:element name="InvoiceRecipientPhoneNumberIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="InvoiceRecipientEmailaddressIdentifier" type="genericStringType0_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="InvoiceRecipientDetailsType">
    <xs:sequence>
      <xs:element name="InvoiceRecipientAddress" type="genericStringType1_35"/>
      <xs:element name="InvoiceRecipientIntermediatorAddress" type="genericNMtokenType8_11"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="InvoiceRecipientPartyDetailsType">
    <xs:sequence>
      <xs:element name="InvoiceRecipientPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="InvoiceRecipientOrganisationName" type="genericStringType2_35" maxOccurs="unbounded"/>
      <xs:element name="InvoiceRecipientDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="InvoiceRecipientOrganisationTaxCode" type="genericNMtokenType0_35" minOccurs="0"/>
      <xs:element name="InvoiceRecipientCode" type="PartyIdentifierType" minOccurs="0"/>
      <xs:element name="InvoiceRecipientPostalAddressDetails" type="InvoiceRecipientPostalAddressDetailsType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="InvoiceRecipientPostalAddressDetailsType">
    <xs:sequence>
      <xs:element name="InvoiceRecipientStreetName" type="genericStringType2_35" maxOccurs="3"/>
      <xs:element name="InvoiceRecipientTownName" type="genericStringType2_35"/>
      <xs:element name="InvoiceRecipientPostCodeIdentifier" type="genericStringType2_35"/>
      <xs:element name="InvoiceRecipientCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="InvoiceRecipientPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:group name="InvoiceRowGroup">
    <xs:sequence>
      <xs:element name="RowSubIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="InvoicedObjectID" type="InvoicedObjectIDType" minOccurs="0"/>
      <xs:element name="ArticleIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="ArticleGroupIdentifier" type="ArticleGroupIdentifierType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="ArticleName" type="genericStringType0_100" minOccurs="0"/>
      <xs:element name="ArticleDescription" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="ArticleInfoUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="BuyerArticleIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="EanCode" type="EanCodeType" minOccurs="0"/>
      <xs:element name="RowRegistrationNumberIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SerialNumberIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowActionCode" type="genericTokenType0_35" minOccurs="0"/>
      <xs:element name="RowDefinitionDetails" type="RowDefinitionDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="OfferedQuantity" type="QuantityType0_14" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="DeliveredQuantity" type="QuantityType0_14" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="OrderedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="ConfirmedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="PostDeliveredQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="InvoicedQuantity" type="QuantityType0_14" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="CreditRequestedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="ReturnedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="StartDate" type="date" minOccurs="0"/>
      <xs:element name="EndDate" type="date" minOccurs="0"/>
      <xs:element name="UnitPriceAmount" type="unitAmountUN" minOccurs="0"/>
      <xs:element name="UnitPriceDiscountAmount" type="unitAmountUN" minOccurs="0"/>
      <xs:element name="UnitPriceNetAmount" type="unitAmountUN" minOccurs="0"/>
      <xs:element name="UnitPriceVatIncludedAmount" type="unitAmountUN" minOccurs="0"/>
      <xs:element name="UnitPriceBaseQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="RowIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="RowOrderPositionIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowIdentifierDate" type="date" minOccurs="0"/>
      <xs:element name="RowPositionIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="OriginalInvoiceNumber" type="genericStringType1_20" minOccurs="0"/>
      <xs:element name="OriginalInvoiceDate" type="date" minOccurs="0"/>
      <xs:element name="OriginalInvoiceReference" type="OriginalInvoiceReferenceType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="RowOrdererName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowSalesPersonName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowOrderConfirmationIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowOrderConfirmationDate" type="date" minOccurs="0"/>
      <xs:element name="RowDeliveryIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowDeliveryIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="RowDeliveryDate" type="date" minOccurs="0"/>
      <xs:element name="RowQuotationIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowQuotationIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="RowAgreementIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowAgreementIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="RowRequestOfQuotationIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowRequestOfQuotationIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="RowPriceListIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowPriceListIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="RowBuyerReferenceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowProjectReferenceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowOverDuePaymentDetails" type="RowOverDuePaymentDetailsType" minOccurs="0"/>
      <xs:element name="RowAnyPartyDetails" type="RowAnyPartyDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="RowDeliveryDetails" type="RowDeliveryDetailsType" minOccurs="0"/>
      <xs:element name="RowShortProposedAccountIdentifier" type="genericNMtokenType0_4" minOccurs="0"/>
      <xs:element name="RowNormalProposedAccountIdentifier" type="genericNMtokenType0_4" minOccurs="0"/>
      <xs:element name="RowProposedAccountText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowAccountDimensionText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowSellerAccountText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowFreeText" type="genericStringType0_512" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="RowUsedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="RowPreviousMeterReadingDate" type="date" minOccurs="0"/>
      <xs:element name="RowLatestMeterReadingDate" type="date" minOccurs="0"/>
      <xs:element name="RowCalculatedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="RowAveragePriceAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowDiscountPercent" type="percentage" minOccurs="0"/>
      <xs:element name="RowDiscountAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowDiscountBaseAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowDiscountTypeCode" type="untdid5189" minOccurs="0"/>
      <xs:element name="RowDiscountTypeText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowProgressiveDiscountDetails" type="RowProgressiveDiscountDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="RowChargeDetails" type="RowChargeDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="RowVatRatePercent" type="percentage" minOccurs="0"/>
      <xs:element name="RowVatCode" type="untdid5305" minOccurs="0"/>
      <xs:element name="RowVatAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowTransactionDetails" type="TransactionDetailsType" minOccurs="0"/>
    </xs:sequence>
  </xs:group>
  <xs:complexType name="SubInvoiceRowType">
    <xs:sequence>
      <xs:element name="SubIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowPositionIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubInvoicedObjectID" type="InvoicedObjectIDType" minOccurs="0"/>
      <xs:element name="SubArticleIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubArticleGroupIdentifier" type="ArticleGroupIdentifierType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SubArticleName" type="genericStringType0_100" minOccurs="0"/>
      <xs:element name="SubArticleDescription" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubArticleInfoUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubBuyerArticleIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubEanCode" type="EanCodeType" minOccurs="0"/>
      <xs:element name="SubRowRegistrationNumberIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubSerialNumberIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowActionCode" type="genericTokenType0_35" minOccurs="0"/>
      <xs:element name="SubRowDefinitionDetails" minOccurs="0" maxOccurs="unbounded">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="SubRowDefinitionHeaderText">
              <xs:complexType>
                <xs:simpleContent>
                  <xs:extension base="genericStringType0_70">
                    <xs:attribute name="DefinitionCode" type="genericTokenType1_20"/>
                  </xs:extension>
                </xs:simpleContent>
              </xs:complexType>
            </xs:element>
            <xs:element name="SubRowDefinitionValue" type="QuantityType0_70" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="SubOfferedQuantity" type="QuantityType0_14" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SubDeliveredQuantity" type="QuantityType0_14" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SubOrderedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubConfirmedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubPostDeliveredQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubInvoicedQuantity" type="QuantityType0_14" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SubCreditRequestedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubReturnedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubStartDate" type="date" minOccurs="0"/>
      <xs:element name="SubEndDate" type="date" minOccurs="0"/>
      <xs:element name="SubUnitPriceAmount" type="unitAmount" minOccurs="0"/>
      <xs:element name="SubUnitPriceDiscountAmount" type="unitAmount" minOccurs="0"/>
      <xs:element name="SubUnitPriceNetAmount" type="unitAmount" minOccurs="0"/>
      <xs:element name="SubUnitPriceVatIncludedAmount" type="unitAmount" minOccurs="0"/>
      <xs:element name="SubUnitPriceBaseQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubRowIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubRowIdentifierDate" type="date" minOccurs="0"/>
      <xs:element name="SubRowOrdererName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowSalesPersonName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowOrderConfirmationIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowOrderConfirmationDate" type="date" minOccurs="0"/>
      <xs:element name="SubOriginalInvoiceNumber" type="genericStringType1_20" minOccurs="0"/>
      <xs:element name="SubOriginalInvoiceDate" type="date" minOccurs="0"/>
      <xs:element name="SubOriginalInvoiceReference" type="OriginalInvoiceReferenceType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SubRowDeliveryIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowDeliveryIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubRowDeliveryDate" type="date" minOccurs="0"/>
      <xs:element name="SubRowQuotationIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowQuotationIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubRowAgreementIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowAgreementIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubRowRequestOfQuotationIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowRequestOfQuotationIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubRowPriceListIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowPriceListIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SubRowBuyerReferenceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowProjectReferenceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowOverDuePaymentDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="SubRowOriginalInvoiceIdentifier" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="SubRowOriginalInvoiceDate" type="date" minOccurs="0"/>
            <xs:element name="SubRowOriginalDueDate" type="date" minOccurs="0"/>
            <xs:element name="SubRowOriginalInvoiceTotalAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowOriginalEpiRemittanceInfoIdentifier" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="SubRowPaidVatExcludedAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowPaidVatIncludedAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowPaidDate" type="date" minOccurs="0"/>
            <xs:element name="SubRowUnPaidVatExcludedAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowUnPaidVatIncludedAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowCollectionDate" type="date" minOccurs="0"/>
            <xs:element name="SubRowCollectionQuantity" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="SubRowCollectionChargeAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowInterestRate" type="percentage" minOccurs="0"/>
            <xs:element name="SubRowInterestStartDate" type="date" minOccurs="0"/>
            <xs:element name="SubRowInterestEndDate" type="date" minOccurs="0"/>
            <xs:element name="SubRowInterestPeriodText" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="SubRowInterestDateNumber" type="genericNMtokenType0_14" minOccurs="0"/>
            <xs:element name="SubRowInterestChargeAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowInterestChargeVatAmount" type="amount" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="SubRowAnyPartyDetails" minOccurs="0" maxOccurs="unbounded">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="SubRowAnyPartyText" type="anypartytexttype0_35"/>
            <xs:element name="SubRowAnyPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
            <xs:element name="SubRowAnyPartyOrganisationName" type="genericStringType2_35" maxOccurs="2"/>
            <xs:element name="SubRowAnyPartyOrganisationDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
            <xs:element name="SubRowAnyPartyOrganisationTaxCode" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="SubRowAnyPartyCode" type="PartyIdentifierType" minOccurs="0"/>
            <xs:element name="SubRowAnyPartyPostalAddressDetails" minOccurs="0">
              <xs:complexType>
                <xs:sequence>
                  <xs:element name="SubRowAnyPartyStreetName" type="genericStringType2_35" maxOccurs="3"/>
                  <xs:element name="SubRowAnyPartyTownName" type="genericStringType2_35"/>
                  <xs:element name="SubRowAnyPartyPostCodeIdentifier" type="genericStringType2_35"/>
                  <xs:element name="SubRowAnyPartyCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
                  <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
                  <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
                  <xs:element name="SubRowAnyPartyPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
                </xs:sequence>
              </xs:complexType>
            </xs:element>
            <xs:element name="SubRowAnyPartyOrganisationUnitNumber" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="SubRowAnyPartySiteCode" type="genericStringType0_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="SubRowDeliveryDetails" type="SubRowDeliveryDetailsType" minOccurs="0"/>
      <xs:element name="SubRowShortProposedAccountIdentifier" type="genericStringType0_4" minOccurs="0"/>
      <xs:element name="SubRowNormalProposedAccountIdentifier" type="genericStringType0_4" minOccurs="0"/>
      <xs:element name="SubRowProposedAccountText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowAccountDimensionText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowSellerAccountText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowFreeText" type="genericStringType0_512" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SubRowUsedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubRowPreviousMeterReadingDate" type="date" minOccurs="0"/>
      <xs:element name="SubRowLatestMeterReadingDate" type="date" minOccurs="0"/>
      <xs:element name="SubRowCalculatedQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="SubRowAveragePriceAmount" type="amount" minOccurs="0"/>
      <xs:element name="SubRowDiscountPercent" type="percentage" minOccurs="0"/>
      <xs:element name="SubRowDiscountAmount" type="amount" minOccurs="0"/>
      <xs:element name="SubRowDiscountBaseAmount" type="amount" minOccurs="0"/>
      <xs:element name="SubRowDiscountTypeCode" type="untdid5189" minOccurs="0"/>
      <xs:element name="SubRowDiscountTypeText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowProgressiveDiscountDetails" minOccurs="0" maxOccurs="unbounded">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="SubRowDiscountPercent" type="percentage" minOccurs="0"/>
            <xs:element name="SubRowDiscountAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowDiscountBaseAmount" type="amount" minOccurs="0"/>
            <xs:element name="SubRowDiscountTypeCode" type="untdid5189" minOccurs="0"/>
            <xs:element name="SubRowDiscountTypeText" type="genericStringType0_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="SubRowChargeDetails" type="RowChargeDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SubRowVatRatePercent" type="percentage" minOccurs="0"/>
      <xs:element name="SubRowVatCode" type="untdid5305" minOccurs="0"/>
      <xs:element name="SubRowVatAmount" type="amount" minOccurs="0"/>
      <xs:element name="SubRowVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="SubRowAmount" type="amount" minOccurs="0"/>
      <xs:element name="SubRowTransactionDetails" type="TransactionDetailsType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="SubRowDeliveryDetailsType">
    <xs:sequence>
      <xs:element name="SubRowTerminalAddressText" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowWaybillIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowWaybillTypeCode" type="genericNMtokenType0_35" minOccurs="0"/>
      <xs:element name="SubRowClearanceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowDeliveryNoteIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowDelivererIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowDelivererName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="SubRowDelivererCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="SubRowDelivererCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="SubRowDelivererCountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowPlaceOfDischarge" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowFinalDestinationName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="SubRowCustomsInfo" type="CustomsInfoType" minOccurs="0"/>
      <xs:element name="SubRowManufacturerArticleIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowManufacturerIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowManufacturerName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="SubRowManufacturerCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="SubRowManufacturerCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="SubRowManufacturerCountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SubRowManufacturerOrderIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SubRowPackageDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="SubRowPackageLength" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="SubRowPackageWidth" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="SubRowPackageHeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="SubRowPackageWeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="SubRowPackageNetWeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="SubRowPackageVolume" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="SubRowTransportCarriageQuantity" type="QuantityType0_14" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="InvoiceRowType">
    <xs:choice>
      <xs:group ref="InvoiceRowGroup"/>
      <xs:element name="SubInvoiceRow" type="SubInvoiceRowType" minOccurs="0" maxOccurs="unbounded"/>
    </xs:choice>
    <!--
    <xs:sequence>
      <xs:element name="SubInvoiceRow" type="SubInvoiceRowType" minOccurs="0" maxOccurs="unbounded"/>
    </xs:sequence>
    -->
  </xs:complexType>
  <xs:complexType name="RowDefinitionDetailsType">
      <xs:sequence>
        <xs:element name="RowDefinitionHeaderText">
          <xs:complexType>
            <xs:simpleContent>
              <xs:extension base="genericStringType0_70">
                <xs:attribute name="DefinitionCode" type="genericTokenType1_20"/>
              </xs:extension>
            </xs:simpleContent>
          </xs:complexType>
        </xs:element>
        <xs:element name="RowDefinitionValue" type="QuantityType0_70" minOccurs="0"/>
      </xs:sequence>
  </xs:complexType>
  <xs:complexType name="RowOverDuePaymentDetailsType">
    <xs:sequence>
      <xs:element name="RowOriginalInvoiceIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowOriginalInvoiceDate" type="date" minOccurs="0"/>
      <xs:element name="RowOriginalDueDate" type="date" minOccurs="0"/>
      <xs:element name="RowOriginalInvoiceTotalAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowOriginalEpiRemittanceInfoIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowPaidVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowPaidVatIncludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowPaidDate" type="date" minOccurs="0"/>
      <xs:element name="RowUnPaidVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowUnPaidVatIncludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowCollectionDate" type="date" minOccurs="0"/>
      <xs:element name="RowCollectionQuantity" type="QuantityType0_14" minOccurs="0"/>
      <xs:element name="RowCollectionChargeAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowInterestRate" type="percentage" minOccurs="0"/>
      <xs:element name="RowInterestStartDate" type="date" minOccurs="0"/>
      <xs:element name="RowInterestEndDate" type="date" minOccurs="0"/>
      <xs:element name="RowInterestPeriodText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowInterestDateNumber" type="genericNMtokenType0_14" minOccurs="0"/>
      <xs:element name="RowInterestChargeAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowInterestChargeVatAmount" type="amount" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="InvoiceSenderPartyDetailsType">
    <xs:sequence>
      <xs:element name="InvoiceSenderPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="InvoiceSenderOrganisationName" type="genericStringType2_35" maxOccurs="unbounded"/>
      <xs:element name="InvoiceSenderOrganisationTaxCode" type="genericNMtokenType0_35" minOccurs="0"/>
      <xs:element name="InvoiceSenderCode" type="PartyIdentifierType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="RowAnyPartyDetailsType">
    <xs:sequence>
      <xs:element name="RowAnyPartyText" type="anypartytexttype0_35"/>
      <xs:element name="RowAnyPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="RowAnyPartyOrganisationName" type="genericStringType2_35" maxOccurs="2"/>
      <xs:element name="RowAnyPartyOrganisationDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="RowAnyPartyOrganisationTaxCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowAnyPartyCode" type="PartyIdentifierType" minOccurs="0"/>
      <xs:element name="RowAnyPartyPostalAddressDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="RowAnyPartyStreetName" type="genericStringType2_35" maxOccurs="3"/>
            <xs:element name="RowAnyPartyTownName" type="genericStringType2_35"/>
            <xs:element name="RowAnyPartyPostCodeIdentifier" type="genericStringType2_35"/>
            <xs:element name="RowAnyPartyCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
            <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
            <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
            <xs:element name="RowAnyPartyPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="RowAnyPartyOrganisationUnitNumber" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowAnyPartySiteCode" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="RowProgressiveDiscountDetailsType">
    <xs:sequence>
      <xs:element name="RowDiscountPercent" type="percentage" minOccurs="0"/>
      <xs:element name="RowDiscountAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowDiscountBaseAmount" type="amount" minOccurs="0"/>
      <xs:element name="RowDiscountTypeCode" type="untdid5189" minOccurs="0"/>
      <xs:element name="RowDiscountTypeText" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:simpleType name="InvoiceTypeCodePatternFI">
    <xs:restriction base="xs:NMTOKEN">
      <xs:pattern value="(REQ|QUO|ORD|ORC|INV|DEV|TES|INF|PRI|DEN|SEI|REC|RES|SDD)[0-9]{2}"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="InvoiceTypeCodeTypeFI">
    <xs:simpleContent>
      <xs:extension base="InvoiceTypeCodePatternFI">
        <xs:attribute name="CodeListAgencyIdentifier">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="SPY"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="InvoiceClassificationType">
    <xs:sequence>
      <xs:element name="ClassificationCode" type="genericStringType1_10" minOccurs="0"/>
      <xs:element name="ClassificationText" type="genericStringType1_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:simpleType name="OriginCodeType">
    <xs:restriction base="xs:NMTOKEN">
      <xs:enumeration value="Original"/>
      <xs:enumeration value="Copy"/>
      <xs:enumeration value="Cancel"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="PartialPaymentDetailsType">
    <xs:sequence>
      <xs:element name="PaidAmount" type="amount"/>
      <xs:element name="PaidVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="UnPaidAmount" type="amount"/>
      <xs:element name="UnPaidVatExcludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="InterestPercent" type="percentage" minOccurs="0"/>
      <xs:element name="ProsessingCostsAmount" type="amount" minOccurs="0"/>
      <xs:element name="PartialPaymentVatIncludedAmount" type="amount" maxOccurs="unbounded"/>
      <xs:element name="PartialPaymentVatExcludedAmount" type="amount" maxOccurs="unbounded"/>
      <xs:element name="PartialPaymentDueDate" type="date" maxOccurs="unbounded"/>
      <xs:element name="PartialPaymentReferenceIdentifier" type="genericStringType2_35" maxOccurs="unbounded"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="PaymentOverDueFineDetailsType">
    <xs:sequence>
      <xs:element name="PaymentOverDueFineFreeText" type="genericStringType0_70" minOccurs="0" maxOccurs="3"/>
      <xs:element name="PaymentOverDueFinePercent" type="percentage" minOccurs="0"/>
      <xs:element name="PaymentOverDueFixedAmount" type="amount" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:simpleType name="PaymentStatusCodeType">
    <xs:restriction base="xs:NMTOKEN">
      <xs:enumeration value="PAID"/>
      <xs:enumeration value="NOTPAID"/>
      <xs:enumeration value="PARTLYPAID"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="PaymentStatusDetailsType">
    <xs:sequence>
      <xs:element name="PaymentStatusCode" type="PaymentStatusCodeType" minOccurs="0"/>
      <xs:element name="PaymentMethodText" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="PaymentTermsDetailsType">
    <xs:sequence>
      <xs:element name="PaymentTermsFreeText" type="genericStringType0_70" minOccurs="0" maxOccurs="2"/>
      <xs:element name="FreeText" type="HeaderValueType" minOccurs="0" maxOccurs="2"/>
      <xs:element name="InvoiceDueDate" type="date" minOccurs="0"/>
      <xs:element name="CashDiscountDate" type="date" minOccurs="0"/>
      <xs:element name="CashDiscountBaseAmount" type="amount" minOccurs="0"/>
      <xs:element name="CashDiscountPercent" type="percentage" minOccurs="0"/>
      <xs:element name="CashDiscountAmount" type="amount" minOccurs="0"/>
      <xs:element name="CashDiscountExcludingVatAmount" type="amount" minOccurs="0"/>
      <xs:element name="CashDiscountVatDetails" minOccurs="0" maxOccurs="unbounded">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="CashDiscountVatPercent" type="percentage"/>
            <xs:element name="CashDiscountVatAmount" type="amount"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="ReducedInvoiceVatIncludedAmount" type="amount" minOccurs="0"/>
      <xs:element name="PaymentOverDueFineDetails" type="PaymentOverDueFineDetailsType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="HeaderValueType">
    <xs:sequence>
      <xs:element name="Header" type="genericStringType1_35" minOccurs="0"/>
      <xs:element name="Value" type="genericStringType1_70" minOccurs="0" maxOccurs="2"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="RowDeliveryDetailsType">
    <xs:sequence>
      <xs:element name="RowTerminalAddressText" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowWaybillIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowWaybillTypeCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowClearanceIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowDeliveryNoteIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowDelivererIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowDelivererName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="RowDelivererCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="RowDelivererCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="RowDelivererCountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowModeOfTransportIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowCarrierName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowVesselName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowLocationIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowTransportInformationDate" type="date" minOccurs="0"/>
      <xs:element name="RowCountryOfOrigin" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowCountryOfDestinationName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowDestinationCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="RowPlaceOfDischarge" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowFinalDestinationName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="RowCustomsInfo" type="CustomsInfoType" minOccurs="0"/>
      <xs:element name="RowManufacturerArticleIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowManufacturerIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowManufacturerName" type="genericStringType0_35" minOccurs="0" maxOccurs="3"/>
      <xs:element name="RowManufacturerCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="RowManufacturerCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="RowManufacturerCountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="RowManufacturerOrderIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="RowPackageDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="RowPackageLength" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="RowPackageWidth" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="RowPackageHeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="RowPackageWeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="RowPackageNetWeight" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="RowPackageVolume" type="QuantityType0_14" minOccurs="0"/>
            <xs:element name="RowTransportCarriageQuantity" type="QuantityType0_14" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="SellerAccountDetailsType">
    <xs:sequence>
      <xs:element name="SellerAccountID" type="SellerAccountIDType"/>
      <xs:element name="SellerBic" type="SellerBicType"/>
      <xs:element name="SellerAccountName" type="genericStringType1_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="SellerAccountIDType">
    <xs:simpleContent>
      <xs:extension base="genericNMtokenType2_35">
        <xs:attribute name="IdentificationSchemeName" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="IBAN"/>
              <xs:enumeration value="BBAN"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="SellerBicType">
    <xs:simpleContent>
      <xs:extension base="genericNMtokenType8_11">
        <xs:attribute name="IdentificationSchemeName" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="BIC"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="SellerCommunicationDetailsType">
    <xs:sequence>
      <xs:element name="SellerPhoneNumberIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerEmailaddressIdentifier" type="genericStringType0_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="SellerInformationDetailsType">
    <xs:sequence>
      <xs:element name="SellerOfficialPostalAddressDetails" minOccurs="0">
        <xs:complexType>
          <xs:sequence>
            <xs:element name="SellerOfficialStreetName" type="genericStringType2_35"/>
            <xs:element name="SellerOfficialTownName" type="genericStringType2_35"/>
            <xs:element name="SellerOfficialPostCodeIdentifier" type="genericStringType2_35"/>
            <xs:element name="SellerOfficialCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
            <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
            <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
      <xs:element name="SellerHomeTownName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerVatRegistrationText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerVatRegistrationDate" type="date" minOccurs="0"/>
      <xs:element name="SellerTaxRegistrationText" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerAdditionalLegalInfo" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SellerPhoneNumber" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerFaxNumber" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerCommonEmailaddressIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SellerWebaddressIdentifier" type="genericStringType0_70" minOccurs="0"/>
      <xs:element name="SellerFreeText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SellerAccountDetails" type="SellerAccountDetailsType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="InvoiceRecipientDetails" type="InvoiceRecipientDetailsType" minOccurs="0" maxOccurs="unbounded"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="SellerPartyDetailsType">
    <xs:sequence>
      <xs:element name="SellerPartyIdentifier" type="PartyLegalRegIdType" minOccurs="0"/>
      <xs:element name="SellerPartyIdentifierUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SellerOrganisationName" type="genericStringType2_70" maxOccurs="unbounded"/>
      <xs:element name="SellerOrganisationTradingName" type="genericStringType2_70" minOccurs="0"/>
      <xs:element name="SellerOrganisationDepartment" type="genericStringType0_35" minOccurs="0" maxOccurs="2"/>
      <xs:element name="SellerOrganisationTaxCode" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerOrganisationTaxCodeUrlText" type="genericStringType0_512" minOccurs="0"/>
      <xs:element name="SellerCode" type="PartyIdentifierType" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="SellerPostalAddressDetails" type="SellerPostalAddressDetailsType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="SellerPostalAddressDetailsType">
    <xs:sequence>
      <xs:element name="SellerStreetName" type="genericStringType2_35" maxOccurs="3"/>
      <xs:element name="SellerTownName" type="genericStringType2_35"/>
      <xs:element name="SellerPostCodeIdentifier" type="genericStringType2_35"/>
      <xs:element name="SellerCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="CountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="CountryName" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="SellerPostOfficeBoxIdentifier" type="genericStringType0_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="AnyPartyCommunicationDetailsType">
    <xs:sequence>
      <xs:element name="AnyPartyPhoneNumberIdentifier" type="genericStringType0_35" minOccurs="0"/>
      <xs:element name="AnyPartyEmailAddressIdentifier" type="genericStringType0_70" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="SpecificationDetailsType">
    <xs:sequence>
      <xs:element name="SpecificationFreeText" type="genericStringType0_80" minOccurs="0" maxOccurs="unbounded"/>
      <xs:element name="ExternalSpecificationDetails" type="ExternalSpecificationDetailsType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="ExternalSpecificationDetailsType">
    <xs:sequence>
      <xs:any namespace="##any" minOccurs="0" maxOccurs="unbounded"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="VatSpecificationDetailsType">
    <xs:sequence>
      <xs:element name="VatBaseAmount" type="amount" minOccurs="0"/>
      <xs:element name="VatRatePercent" type="percentage" minOccurs="0"/>
      <xs:element name="VatCode" type="untdid5305" minOccurs="0"/>
      <xs:element name="VatRateAmount" type="amount" minOccurs="0"/>
      <xs:element name="VatFreeText" type="genericStringType0_70" minOccurs="0" maxOccurs="3"/>
      <xs:element name="VatExemptionReasonCode" type="VatExReasonCodeType" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:simpleType name="VatExReasonCodeType">
    <xs:restriction base="xs:NMTOKEN">
      <xs:pattern value="[A-Za-z0-9\-]{1,20}"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="VatPointType">
    <xs:choice>
      <xs:element name="VatPointDate" type="date"/>
      <xs:element name="VatPointDateCode" type="untdid2005"/>
    </xs:choice>
  </xs:complexType>
  <xs:complexType name="PartyIdentifierType">
    <xs:simpleContent>
      <xs:extension base="genericStringType0_70">
        <xs:attribute name="IdentifierType" type="genericTokenType1_20"/>
        <xs:attribute name="SchemeID" type="iso6523cid"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="PartyLegalRegIdType">
    <xs:simpleContent>
      <xs:extension base="genericStringType0_35">
        <xs:attribute name="SchemeID" type="iso6523cid"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="VatCategoryCodeType">
      <xs:restriction base="untdid5305"/>
  </xs:simpleType>
  <xs:complexType name="DiscountDetailsType">
    <xs:sequence>
      <xs:element name="FreeText" type="genericStringType1_70" minOccurs="0"/>
      <xs:element name="ReasonCode" type="untdid5189" minOccurs="0"/>
      <xs:element name="Percent" type="percentage" minOccurs="0"/>
      <xs:element name="Amount" type="amount" minOccurs="0"/>
      <xs:element name="BaseAmount" type="amount" minOccurs="0"/>
      <xs:element name="VatCategoryCode" type ="VatCategoryCodeType" minOccurs="0"/>
      <xs:element name="VatRatePercent" type="percentage" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="InvoiceChargeDetailsType">
    <xs:sequence>
      <xs:element name="ReasonText" type="genericStringType1_70" minOccurs="0"/>
      <xs:element name="ReasonCode" type="untdid7161" minOccurs="0"/>
      <xs:element name="Percent" type="percentage" minOccurs="0"/>
      <xs:element name="Amount" type="amount" minOccurs="0"/>
      <xs:element name="BaseAmount" type="amount" minOccurs="0"/>
      <xs:element name="VatCategoryCode" type ="VatCategoryCodeType" minOccurs="0"/>
      <xs:element name="VatRatePercent" type="percentage" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="RowChargeDetailsType">
    <xs:sequence>
      <xs:element name="ReasonText" type="genericStringType1_70" minOccurs="0"/>
      <xs:element name="ReasonCode" type="untdid7161" minOccurs="0"/>
      <xs:element name="Percent" type="percentage" minOccurs="0"/>
      <xs:element name="Amount" type="amount" minOccurs="0"/>
      <xs:element name="BaseAmount" type="amount" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="CustomsInfoType">
    <xs:sequence>
      <xs:element name="CNCode" type="genericStringType1_8" minOccurs="0"/>
      <xs:element name="CNName" type="genericStringType1_35" minOccurs="0"/>
      <xs:element name="CNOriginCountrySubdivision" type="genericStringType2_35" minOccurs="0"/>
      <xs:element name="CNOriginCountryCode" type="CountryCodeType" minOccurs="0"/>
      <xs:element name="CNOriginCountryName" type="genericStringType1_35" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="TransactionDetailsType">
    <xs:sequence>
      <xs:element name="OtherCurrencyAmount" type="amount" minOccurs="0"/>
      <xs:element name="ExchangeRate" type="exchangeRate" minOccurs="0"/>
      <xs:element name="ExchangeDate" type="date" minOccurs="0"/>
    </xs:sequence>
  </xs:complexType>
  <xs:complexType name="AttachmentMessageDetailsType">
    <xs:sequence>
      <xs:element name="AttachmentMessageIdentifier" type="AttachmentsIdentifierType"/>
    </xs:sequence>
  </xs:complexType>
  <!-- Common types -->
  <xs:complexType name="QuantityType">
    <xs:simpleContent>
      <xs:extension base="xs:string">
        <xs:attribute name="QuantityUnitCode" type="genericTokenType0_14"/>
        <xs:attribute name="QuantityUnitCodeUN" type="unitCodeUN"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="QuantityType0_14">
    <xs:simpleContent>
      <xs:restriction base="QuantityType">
        <xs:minLength value="0"/>
        <xs:maxLength value="14"/>
      </xs:restriction>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="QuantityType0_70">
    <xs:simpleContent>
      <xs:restriction base="QuantityType">
        <xs:minLength value="0"/>
        <xs:maxLength value="70"/>
      </xs:restriction>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="AnyPartyTextType">
    <xs:simpleContent>
      <xs:extension base="xs:string">
        <xs:attribute name="AnyPartyCode" type="genericTokenType0_35" use="required"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="anypartytexttype0_35">
    <xs:simpleContent>
      <xs:restriction base="AnyPartyTextType">
        <xs:minLength value="0"/>
        <xs:maxLength value="35"/>
      </xs:restriction>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="AttachmentsIdentifierType">
    <xs:restriction base="xs:string">
      <xs:minLength value="15"/>
      <xs:maxLength value="61"/>
      <xs:pattern value=".{2,48}::attachments"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericTokenType3">
    <xs:restriction base="xs:token">
      <xs:length value="3"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericTokenType0_14">
    <xs:restriction base="xs:token">
      <xs:minLength value="0"/>
      <xs:maxLength value="14"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericTokenType1_20">
    <xs:restriction base="xs:token">
      <xs:minLength value="1"/>
      <xs:maxLength value="20"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericTokenType0_35">
    <xs:restriction base="xs:token">
      <xs:minLength value="0"/>
      <xs:maxLength value="35"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericTokenType2_35">
    <xs:restriction base="xs:token">
      <xs:minLength value="2"/>
      <xs:maxLength value="35"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericTokenType0_70">
    <xs:restriction base="xs:token">
      <xs:minLength value="0"/>
      <xs:maxLength value="70"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericNMtokenType0_4">
    <xs:restriction base="xs:token">
      <xs:minLength value="0"/>
      <xs:maxLength value="4"/>
      <xs:pattern value="\c*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericNMtokenType8_11">
    <xs:restriction base="xs:token">
      <xs:minLength value="8"/>
      <xs:maxLength value="11"/>
      <xs:pattern value="\c*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericNMtokenType0_14">
    <xs:restriction base="xs:token">
      <xs:minLength value="0"/>
      <xs:maxLength value="14"/>
      <xs:pattern value="\c*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericNMtokenType1_34">
    <xs:restriction base="xs:token">
      <xs:minLength value="1"/>
      <xs:maxLength value="34"/>
      <xs:pattern value="\c*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericNMtokenType0_35">
    <xs:restriction base="xs:token">
      <xs:minLength value="0"/>
      <xs:maxLength value="35"/>
      <xs:pattern value="\c*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericNMtokenType2_35">
    <xs:restriction base="xs:token">
      <xs:minLength value="2"/>
      <xs:maxLength value="35"/>
      <xs:pattern value="\c*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericNMtokenType0_512">
    <xs:restriction base="xs:token">
      <xs:minLength value="0"/>
      <xs:maxLength value="512"/>
      <xs:pattern value="\c*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_4">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="4"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_4">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="4"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_8">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="8"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_10">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="10"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_10">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="10"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_14">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="14"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_13">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="13"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_19">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="19"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_20">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="20"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_35">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="35"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_35">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="35"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType2_35">
    <xs:restriction base="xs:string">
      <xs:minLength value="2"/>
      <xs:maxLength value="35"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_48">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="48"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType2_48">
    <xs:restriction base="xs:string">
      <xs:minLength value="2"/>
      <xs:maxLength value="48"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_70">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="70"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType1_70">
    <xs:restriction base="xs:string">
      <xs:minLength value="1"/>
      <xs:maxLength value="70"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType2_70">
    <xs:restriction base="xs:string">
      <xs:minLength value="2"/>
      <xs:maxLength value="70"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_80">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="80"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_100">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="100"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="genericStringType0_512">
    <xs:restriction base="xs:string">
      <xs:minLength value="0"/>
      <xs:maxLength value="512"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="percentage">
    <xs:restriction base="xs:string">
      <xs:pattern value="[1-9]?[0-9]{1,2}(,[0-9]{1,4})?"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="exchangeRate">
    <xs:restriction base="xs:string">
      <xs:pattern value="[0-9]{1,15}(,[0-9]{1,6})?"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="date">
    <xs:simpleContent>
      <xs:extension base="dateType">
        <xs:attribute name="Format" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:enumeration value="CCYYMMDD"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="dateType">
    <xs:restriction base="xs:integer">
      <xs:pattern value="[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="amount">
    <xs:simpleContent>
      <xs:extension base="monetaryAmount">
        <xs:attribute name="AmountCurrencyIdentifier" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:length value="3"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="monetaryAmount">
    <xs:restriction base="xs:token">
      <xs:pattern value="-?[0-9]{1,15}(,[0-9]{2,5})?"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="epiAmount">
    <xs:simpleContent>
      <xs:extension base="epiMonetaryAmount">
        <xs:attribute name="AmountCurrencyIdentifier" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:length value="3"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="epiMonetaryAmount">
    <xs:restriction base="xs:token">
      <xs:pattern value="-?[0-9]{1,15},[0-9]{2}"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="unitAmount">
    <xs:simpleContent>
      <xs:extension base="unitAmountType">
        <xs:attribute name="AmountCurrencyIdentifier" use="required">
          <xs:simpleType>
            <xs:restriction base="xs:NMTOKEN">
              <xs:length value="3"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:attribute>
        <xs:attribute name="UnitPriceUnitCode" type="genericStringType0_14"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="unitAmountUN">
    <xs:simpleContent>
      <xs:extension base="unitAmount">
        <xs:attribute name="QuantityUnitCodeUN" type="unitCodeUN"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="unitAmountType">
    <xs:restriction base="xs:token">
      <xs:pattern value="-?[0-9]{1,15}(,[0-9]{2,5})?"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="CountryCodeType">
    <xs:restriction base="xs:NMTOKEN">
      <xs:length value="2"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="LanguageCodeType">
    <xs:restriction base="xs:NMTOKEN">
      <xs:length value="2"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="iso6523cid">
    <xs:restriction base="xs:NMTOKEN">
      <xs:pattern value="[0-9]{4}"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="EanCodeType">
    <xs:simpleContent>
      <xs:extension base="genericStringType0_35">
        <xs:attribute name="SchemeID" type="iso6523cid"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="DestinationNameType">
    <xs:simpleContent>
      <xs:extension base="genericStringType0_35">
        <xs:attribute name="SchemeID" type="iso6523cid"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="ArticleGroupIdentifierType">
    <xs:simpleContent>
      <xs:extension base="genericStringType0_70">
        <xs:attribute name="SchemeID" type="untdid7143"/>
        <xs:attribute name="SchemeVersion" type="genericStringType1_35"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:complexType name="InvoicedObjectIDType">
    <xs:simpleContent>
      <xs:extension base="genericStringType1_70">
        <xs:attribute name="SchemeID" type="untdid1153"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="capAZ09">
    <xs:restriction base="xs:NMTOKEN">
      <xs:pattern value="[A-Z0-9]*"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="untdid1_3_AZ09">
    <xs:restriction base="capAZ09">
      <xs:minLength value="1"/>
      <xs:maxLength value="3"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="untdid1001">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="untdid1153">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="untdid2005">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="untdid4461">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="untdid5189">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="untdid5305">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="untdid7143">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="untdid7161">
    <xs:restriction base="untdid1_3_AZ09"/>
  </xs:simpleType>
  <xs:simpleType name="ElectronicAddrSchemeIdType">
    <xs:restriction base="xs:NMTOKEN">
      <xs:minLength value="1"/>
      <xs:maxLength value="10"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:complexType name="ElectronicAddrIdType">
    <xs:simpleContent>
      <xs:extension base="genericStringType2_35">
        <xs:attribute name="SchemeID" type="ElectronicAddrSchemeIdType"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>
  <xs:simpleType name="unitCodeUN">
    <xs:restriction base="xs:NMTOKEN">
      <xs:pattern value="[A-Z0-9]{2,3}"/>
    </xs:restriction>
  </xs:simpleType>
</xs:schema>`)
