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
)

// checks if the TokenClaims type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenClaims{}

// TokenClaims struct for TokenClaims
type TokenClaims struct {
	Aud *string `json:"aud,omitempty"`
	Email *string `json:"email,omitempty"`
	Exp *string `json:"exp,omitempty"`
	Extra *map[string][]string `json:"extra,omitempty"`
	Iat *string `json:"iat,omitempty"`
	Iss *string `json:"iss,omitempty"`
	Jti *string `json:"jti,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Nbf *string `json:"nbf,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PreferredUsername *string `json:"preferred_username,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Sub *string `json:"sub,omitempty"`
	TokenType *string `json:"token_type,omitempty"`
	Url *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewTokenClaims instantiates a new TokenClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenClaims() *TokenClaims {
	this := TokenClaims{}
	return &this
}

// NewTokenClaimsWithDefaults instantiates a new TokenClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenClaimsWithDefaults() *TokenClaims {
	this := TokenClaims{}
	return &this
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *TokenClaims) GetAud() string {
	if o == nil || IsNil(o.Aud) {
		var ret string
		return ret
	}
	return *o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetAudOk() (*string, bool) {
	if o == nil || IsNil(o.Aud) {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *TokenClaims) HasAud() bool {
	if o != nil && !IsNil(o.Aud) {
		return true
	}

	return false
}

// SetAud gets a reference to the given string and assigns it to the Aud field.
func (o *TokenClaims) SetAud(v string) {
	o.Aud = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TokenClaims) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TokenClaims) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *TokenClaims) SetEmail(v string) {
	o.Email = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *TokenClaims) GetExp() string {
	if o == nil || IsNil(o.Exp) {
		var ret string
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetExpOk() (*string, bool) {
	if o == nil || IsNil(o.Exp) {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *TokenClaims) HasExp() bool {
	if o != nil && !IsNil(o.Exp) {
		return true
	}

	return false
}

// SetExp gets a reference to the given string and assigns it to the Exp field.
func (o *TokenClaims) SetExp(v string) {
	o.Exp = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *TokenClaims) GetExtra() map[string][]string {
	if o == nil || IsNil(o.Extra) {
		var ret map[string][]string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetExtraOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *TokenClaims) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string][]string and assigns it to the Extra field.
func (o *TokenClaims) SetExtra(v map[string][]string) {
	o.Extra = &v
}

// GetIat returns the Iat field value if set, zero value otherwise.
func (o *TokenClaims) GetIat() string {
	if o == nil || IsNil(o.Iat) {
		var ret string
		return ret
	}
	return *o.Iat
}

// GetIatOk returns a tuple with the Iat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetIatOk() (*string, bool) {
	if o == nil || IsNil(o.Iat) {
		return nil, false
	}
	return o.Iat, true
}

// HasIat returns a boolean if a field has been set.
func (o *TokenClaims) HasIat() bool {
	if o != nil && !IsNil(o.Iat) {
		return true
	}

	return false
}

// SetIat gets a reference to the given string and assigns it to the Iat field.
func (o *TokenClaims) SetIat(v string) {
	o.Iat = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *TokenClaims) GetIss() string {
	if o == nil || IsNil(o.Iss) {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetIssOk() (*string, bool) {
	if o == nil || IsNil(o.Iss) {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *TokenClaims) HasIss() bool {
	if o != nil && !IsNil(o.Iss) {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *TokenClaims) SetIss(v string) {
	o.Iss = &v
}

// GetJti returns the Jti field value if set, zero value otherwise.
func (o *TokenClaims) GetJti() string {
	if o == nil || IsNil(o.Jti) {
		var ret string
		return ret
	}
	return *o.Jti
}

// GetJtiOk returns a tuple with the Jti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetJtiOk() (*string, bool) {
	if o == nil || IsNil(o.Jti) {
		return nil, false
	}
	return o.Jti, true
}

// HasJti returns a boolean if a field has been set.
func (o *TokenClaims) HasJti() bool {
	if o != nil && !IsNil(o.Jti) {
		return true
	}

	return false
}

// SetJti gets a reference to the given string and assigns it to the Jti field.
func (o *TokenClaims) SetJti(v string) {
	o.Jti = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *TokenClaims) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *TokenClaims) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *TokenClaims) SetLocale(v string) {
	o.Locale = &v
}

// GetNbf returns the Nbf field value if set, zero value otherwise.
func (o *TokenClaims) GetNbf() string {
	if o == nil || IsNil(o.Nbf) {
		var ret string
		return ret
	}
	return *o.Nbf
}

// GetNbfOk returns a tuple with the Nbf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetNbfOk() (*string, bool) {
	if o == nil || IsNil(o.Nbf) {
		return nil, false
	}
	return o.Nbf, true
}

// HasNbf returns a boolean if a field has been set.
func (o *TokenClaims) HasNbf() bool {
	if o != nil && !IsNil(o.Nbf) {
		return true
	}

	return false
}

// SetNbf gets a reference to the given string and assigns it to the Nbf field.
func (o *TokenClaims) SetNbf(v string) {
	o.Nbf = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *TokenClaims) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *TokenClaims) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *TokenClaims) SetNonce(v string) {
	o.Nonce = &v
}

// GetPreferredUsername returns the PreferredUsername field value if set, zero value otherwise.
func (o *TokenClaims) GetPreferredUsername() string {
	if o == nil || IsNil(o.PreferredUsername) {
		var ret string
		return ret
	}
	return *o.PreferredUsername
}

// GetPreferredUsernameOk returns a tuple with the PreferredUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetPreferredUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredUsername) {
		return nil, false
	}
	return o.PreferredUsername, true
}

// HasPreferredUsername returns a boolean if a field has been set.
func (o *TokenClaims) HasPreferredUsername() bool {
	if o != nil && !IsNil(o.PreferredUsername) {
		return true
	}

	return false
}

// SetPreferredUsername gets a reference to the given string and assigns it to the PreferredUsername field.
func (o *TokenClaims) SetPreferredUsername(v string) {
	o.PreferredUsername = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenClaims) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenClaims) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenClaims) SetScopes(v []string) {
	o.Scopes = v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *TokenClaims) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *TokenClaims) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *TokenClaims) SetSub(v string) {
	o.Sub = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *TokenClaims) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenClaims) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *TokenClaims) SetTokenType(v string) {
	o.TokenType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TokenClaims) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TokenClaims) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TokenClaims) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TokenClaims) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenClaims) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TokenClaims) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TokenClaims) SetUsername(v string) {
	o.Username = &v
}

func (o TokenClaims) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenClaims) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aud) {
		toSerialize["aud"] = o.Aud
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Exp) {
		toSerialize["exp"] = o.Exp
	}
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !IsNil(o.Iat) {
		toSerialize["iat"] = o.Iat
	}
	if !IsNil(o.Iss) {
		toSerialize["iss"] = o.Iss
	}
	if !IsNil(o.Jti) {
		toSerialize["jti"] = o.Jti
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Nbf) {
		toSerialize["nbf"] = o.Nbf
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.PreferredUsername) {
		toSerialize["preferred_username"] = o.PreferredUsername
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableTokenClaims struct {
	value *TokenClaims
	isSet bool
}

func (v NullableTokenClaims) Get() *TokenClaims {
	return v.value
}

func (v *NullableTokenClaims) Set(val *TokenClaims) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenClaims) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenClaims(val *TokenClaims) *NullableTokenClaims {
	return &NullableTokenClaims{value: val, isSet: true}
}

func (v NullableTokenClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


