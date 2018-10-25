# io.pharmatrace.api.scin.public\ContainerApi

All URIs are relative to *https://api.pharmatrace.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationHistoryGet**](ContainerApi.md#LocationHistoryGet) | **Get** /locationHistory | Location History


# **LocationHistoryGet**
> Locations LocationHistoryGet(ctx, optional)
Location History

Returns the location history of a container or a individually labeled pharmaceutical

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationHistoryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationHistoryGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerId** | **optional.String**| container id of the tracked container. | 
 **offset** | **optional.Int32**| Offset the list of returned results by this amount. Default is zero. | 
 **limit** | **optional.Int32**| Number of items to retrieve. Default is 5, maximum is 100. | 

### Return type

[**Locations**](Locations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

