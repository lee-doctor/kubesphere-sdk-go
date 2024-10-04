/*
KS API

KubeSphere OpenAPI

API version: v4.1.1
Contact: support@kubesphere.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OauthProviderMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OauthProviderMetadata{}

// OauthProviderMetadata struct for OauthProviderMetadata
type OauthProviderMetadata struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	ClaimsSupported []string `json:"claims_supported"`
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported"`
	GrantTypesSupported []string `json:"grant_types_supported"`
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	Issuer string `json:"issuer"`
	JwksUri string `json:"jwks_uri"`
	ResponseTypesSupported []string `json:"response_types_supported"`
	ScopesSupported []string `json:"scopes_supported"`
	SubjectTypesSupported []string `json:"subject_types_supported"`
	TokenEndpoint string `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	UserinfoEndpoint string `json:"userinfo_endpoint"`
}

type _OauthProviderMetadata OauthProviderMetadata

// NewOauthProviderMetadata instantiates a new OauthProviderMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthProviderMetadata(authorizationEndpoint string, claimsSupported []string, codeChallengeMethodsSupported []string, grantTypesSupported []string, idTokenSigningAlgValuesSupported []string, issuer string, jwksUri string, responseTypesSupported []string, scopesSupported []string, subjectTypesSupported []string, tokenEndpoint string, tokenEndpointAuthMethodsSupported []string, userinfoEndpoint string) *OauthProviderMetadata {
	this := OauthProviderMetadata{}
	this.AuthorizationEndpoint = authorizationEndpoint
	this.ClaimsSupported = claimsSupported
	this.CodeChallengeMethodsSupported = codeChallengeMethodsSupported
	this.GrantTypesSupported = grantTypesSupported
	this.IdTokenSigningAlgValuesSupported = idTokenSigningAlgValuesSupported
	this.Issuer = issuer
	this.JwksUri = jwksUri
	this.ResponseTypesSupported = responseTypesSupported
	this.ScopesSupported = scopesSupported
	this.SubjectTypesSupported = subjectTypesSupported
	this.TokenEndpoint = tokenEndpoint
	this.TokenEndpointAuthMethodsSupported = tokenEndpointAuthMethodsSupported
	this.UserinfoEndpoint = userinfoEndpoint
	return &this
}

// NewOauthProviderMetadataWithDefaults instantiates a new OauthProviderMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthProviderMetadataWithDefaults() *OauthProviderMetadata {
	this := OauthProviderMetadata{}
	return &this
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value
func (o *OauthProviderMetadata) GetAuthorizationEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationEndpoint, true
}

// SetAuthorizationEndpoint sets field value
func (o *OauthProviderMetadata) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = v
}

// GetClaimsSupported returns the ClaimsSupported field value
func (o *OauthProviderMetadata) GetClaimsSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ClaimsSupported
}

// GetClaimsSupportedOk returns a tuple with the ClaimsSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetClaimsSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClaimsSupported, true
}

// SetClaimsSupported sets field value
func (o *OauthProviderMetadata) SetClaimsSupported(v []string) {
	o.ClaimsSupported = v
}

// GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field value
func (o *OauthProviderMetadata) GetCodeChallengeMethodsSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CodeChallengeMethodsSupported
}

// GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetCodeChallengeMethodsSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeChallengeMethodsSupported, true
}

// SetCodeChallengeMethodsSupported sets field value
func (o *OauthProviderMetadata) SetCodeChallengeMethodsSupported(v []string) {
	o.CodeChallengeMethodsSupported = v
}

// GetGrantTypesSupported returns the GrantTypesSupported field value
func (o *OauthProviderMetadata) GetGrantTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GrantTypesSupported
}

// GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetGrantTypesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypesSupported, true
}

// SetGrantTypesSupported sets field value
func (o *OauthProviderMetadata) SetGrantTypesSupported(v []string) {
	o.GrantTypesSupported = v
}

// GetIdTokenSigningAlgValuesSupported returns the IdTokenSigningAlgValuesSupported field value
func (o *OauthProviderMetadata) GetIdTokenSigningAlgValuesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IdTokenSigningAlgValuesSupported
}

// GetIdTokenSigningAlgValuesSupportedOk returns a tuple with the IdTokenSigningAlgValuesSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetIdTokenSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdTokenSigningAlgValuesSupported, true
}

// SetIdTokenSigningAlgValuesSupported sets field value
func (o *OauthProviderMetadata) SetIdTokenSigningAlgValuesSupported(v []string) {
	o.IdTokenSigningAlgValuesSupported = v
}

// GetIssuer returns the Issuer field value
func (o *OauthProviderMetadata) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OauthProviderMetadata) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwksUri returns the JwksUri field value
func (o *OauthProviderMetadata) GetJwksUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetJwksUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwksUri, true
}

// SetJwksUri sets field value
func (o *OauthProviderMetadata) SetJwksUri(v string) {
	o.JwksUri = v
}

// GetResponseTypesSupported returns the ResponseTypesSupported field value
func (o *OauthProviderMetadata) GetResponseTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResponseTypesSupported
}

// GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetResponseTypesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseTypesSupported, true
}

// SetResponseTypesSupported sets field value
func (o *OauthProviderMetadata) SetResponseTypesSupported(v []string) {
	o.ResponseTypesSupported = v
}

// GetScopesSupported returns the ScopesSupported field value
func (o *OauthProviderMetadata) GetScopesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScopesSupported
}

// GetScopesSupportedOk returns a tuple with the ScopesSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetScopesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopesSupported, true
}

// SetScopesSupported sets field value
func (o *OauthProviderMetadata) SetScopesSupported(v []string) {
	o.ScopesSupported = v
}

// GetSubjectTypesSupported returns the SubjectTypesSupported field value
func (o *OauthProviderMetadata) GetSubjectTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubjectTypesSupported
}

// GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetSubjectTypesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypesSupported, true
}

// SetSubjectTypesSupported sets field value
func (o *OauthProviderMetadata) SetSubjectTypesSupported(v []string) {
	o.SubjectTypesSupported = v
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *OauthProviderMetadata) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *OauthProviderMetadata) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field value
func (o *OauthProviderMetadata) GetTokenEndpointAuthMethodsSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TokenEndpointAuthMethodsSupported
}

// GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetTokenEndpointAuthMethodsSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethodsSupported, true
}

// SetTokenEndpointAuthMethodsSupported sets field value
func (o *OauthProviderMetadata) SetTokenEndpointAuthMethodsSupported(v []string) {
	o.TokenEndpointAuthMethodsSupported = v
}

// GetUserinfoEndpoint returns the UserinfoEndpoint field value
func (o *OauthProviderMetadata) GetUserinfoEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserinfoEndpoint
}

// GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field value
// and a boolean to check if the value has been set.
func (o *OauthProviderMetadata) GetUserinfoEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserinfoEndpoint, true
}

// SetUserinfoEndpoint sets field value
func (o *OauthProviderMetadata) SetUserinfoEndpoint(v string) {
	o.UserinfoEndpoint = v
}

func (o OauthProviderMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OauthProviderMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorization_endpoint"] = o.AuthorizationEndpoint
	toSerialize["claims_supported"] = o.ClaimsSupported
	toSerialize["code_challenge_methods_supported"] = o.CodeChallengeMethodsSupported
	toSerialize["grant_types_supported"] = o.GrantTypesSupported
	toSerialize["id_token_signing_alg_values_supported"] = o.IdTokenSigningAlgValuesSupported
	toSerialize["issuer"] = o.Issuer
	toSerialize["jwks_uri"] = o.JwksUri
	toSerialize["response_types_supported"] = o.ResponseTypesSupported
	toSerialize["scopes_supported"] = o.ScopesSupported
	toSerialize["subject_types_supported"] = o.SubjectTypesSupported
	toSerialize["token_endpoint"] = o.TokenEndpoint
	toSerialize["token_endpoint_auth_methods_supported"] = o.TokenEndpointAuthMethodsSupported
	toSerialize["userinfo_endpoint"] = o.UserinfoEndpoint
	return toSerialize, nil
}

func (o *OauthProviderMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorization_endpoint",
		"claims_supported",
		"code_challenge_methods_supported",
		"grant_types_supported",
		"id_token_signing_alg_values_supported",
		"issuer",
		"jwks_uri",
		"response_types_supported",
		"scopes_supported",
		"subject_types_supported",
		"token_endpoint",
		"token_endpoint_auth_methods_supported",
		"userinfo_endpoint",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOauthProviderMetadata := _OauthProviderMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOauthProviderMetadata)

	if err != nil {
		return err
	}

	*o = OauthProviderMetadata(varOauthProviderMetadata)

	return err
}

type NullableOauthProviderMetadata struct {
	value *OauthProviderMetadata
	isSet bool
}

func (v NullableOauthProviderMetadata) Get() *OauthProviderMetadata {
	return v.value
}

func (v *NullableOauthProviderMetadata) Set(val *OauthProviderMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthProviderMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthProviderMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthProviderMetadata(val *OauthProviderMetadata) *NullableOauthProviderMetadata {
	return &NullableOauthProviderMetadata{value: val, isSet: true}
}

func (v NullableOauthProviderMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthProviderMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

