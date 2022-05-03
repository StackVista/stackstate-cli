# TopologyMatchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchedCheckStates** | **int32** |  | 
**UnmatchedCheckStates** | [**[]UnmatchedCheckState**](UnmatchedCheckState.md) |  | 
**MultipleMatchesCheckStates** | [**[]MultipleMatchesCheckState**](MultipleMatchesCheckState.md) |  | 

## Methods

### NewTopologyMatchResult

`func NewTopologyMatchResult(matchedCheckStates int32, unmatchedCheckStates []UnmatchedCheckState, multipleMatchesCheckStates []MultipleMatchesCheckState, ) *TopologyMatchResult`

NewTopologyMatchResult instantiates a new TopologyMatchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyMatchResultWithDefaults

`func NewTopologyMatchResultWithDefaults() *TopologyMatchResult`

NewTopologyMatchResultWithDefaults instantiates a new TopologyMatchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchedCheckStates

`func (o *TopologyMatchResult) GetMatchedCheckStates() int32`

GetMatchedCheckStates returns the MatchedCheckStates field if non-nil, zero value otherwise.

### GetMatchedCheckStatesOk

`func (o *TopologyMatchResult) GetMatchedCheckStatesOk() (*int32, bool)`

GetMatchedCheckStatesOk returns a tuple with the MatchedCheckStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCheckStates

`func (o *TopologyMatchResult) SetMatchedCheckStates(v int32)`

SetMatchedCheckStates sets MatchedCheckStates field to given value.


### GetUnmatchedCheckStates

`func (o *TopologyMatchResult) GetUnmatchedCheckStates() []UnmatchedCheckState`

GetUnmatchedCheckStates returns the UnmatchedCheckStates field if non-nil, zero value otherwise.

### GetUnmatchedCheckStatesOk

`func (o *TopologyMatchResult) GetUnmatchedCheckStatesOk() (*[]UnmatchedCheckState, bool)`

GetUnmatchedCheckStatesOk returns a tuple with the UnmatchedCheckStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedCheckStates

`func (o *TopologyMatchResult) SetUnmatchedCheckStates(v []UnmatchedCheckState)`

SetUnmatchedCheckStates sets UnmatchedCheckStates field to given value.


### GetMultipleMatchesCheckStates

`func (o *TopologyMatchResult) GetMultipleMatchesCheckStates() []MultipleMatchesCheckState`

GetMultipleMatchesCheckStates returns the MultipleMatchesCheckStates field if non-nil, zero value otherwise.

### GetMultipleMatchesCheckStatesOk

`func (o *TopologyMatchResult) GetMultipleMatchesCheckStatesOk() (*[]MultipleMatchesCheckState, bool)`

GetMultipleMatchesCheckStatesOk returns a tuple with the MultipleMatchesCheckStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleMatchesCheckStates

`func (o *TopologyMatchResult) SetMultipleMatchesCheckStates(v []MultipleMatchesCheckState)`

SetMultipleMatchesCheckStates sets MultipleMatchesCheckStates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


