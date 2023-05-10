package google

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var textListType = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: `A collection of text responses.`,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
}
var boolType = &schema.Schema{
	Type:        schema.TypeBool,
	Computed:    true,
	Description: `Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.`,
}
var messagesTextType = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: `The text response message.`,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"text":                        textListType,
			"allow_playback_interruption": boolType,
		},
	},
}
var payloadImageType = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"src": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
			"alt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
		},
	},
}
var payloadSimpleListType = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 10,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"icon": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
			"event": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
		},
	},
}
var payloadQuickRepliesType = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 3,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
		},
	},
}
var payloadListSuggestionType = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 10,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"icon": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
			"event": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A collection of text responses.`,
			},
		},
	},
}
var messagesPayloadType = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: `Payload responses`,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"text":             textListType,
			"image":            payloadImageType,
			"simple_list":      payloadSimpleListType,
			"quick_replies":    payloadQuickRepliesType,
			"list_suggestions": payloadListSuggestionType,
		},
	},
}
var messagesType = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: `The list of rich message responses to present to the user.`,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"text":    messagesTextType,
			"payload": messagesPayloadType,
		},
	},
}
var setParameterActionsType = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: `Set Parameters`,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parameter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.`,
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.`,
			},
		},
	},
}
var fulfillmentType = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: `The fulfillment to call when the condition is satisfied. At least one of triggerFulfillment and target must be specified. When both are defined, triggerFulfillment is executed first.`,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"messages":              messagesType,
			"set_parameter_actions": setParameterActionsType,
			"return_partial_responses": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Whether Dialogflow should return currently queued fulfillment response messages in streaming APIs. If a webhook is specified, it happens before Dialogflow invokes webhook. Warning: 1) This flag only affects streaming API. Responses are still queued and returned once in non-streaming API. 2) The flag can be enabled in any fulfillment but only the first 3 partial responses will be returned. You may only want to apply it to fulfillments that have slow webhooks.`,
			},
			"tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The tag used by the webhook to identify which fulfillment is being called. This field is required if webhook is specified.`,
			},
			"webhook": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The webhook to call. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.`,
			},
		},
	},
}
var eventHandlersType = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: `Handlers associated with the page to handle events such as webhook errors, no match or no input.`,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The name of the event to handle.`,
			},
			"target_flow": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The target flow to transition to.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.`,
			},
			"target_page": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The target page to transition to.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.`,
			},
			"trigger_fulfillment": fulfillmentType,
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier of this event handler.`,
			},
		},
	},
}

func flattenDialogflowCXFulfillment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["messages"] =
		flattenDialogflowCXFulfillmentMessages(original["messages"], d, config)
	transformed["set_parameter_actions"] = original["setParameterActions"]
	transformed["webhook"] = original["webhook"]
	// flattenDialogflowCXFulfillmentWebhook(original["webhook"], d, config)
	transformed["return_partial_responses"] = original["return_partial_responses"]
	// flattenDialogflowCXFulfillmentReturnPartialResponses(original["returnPartialResponses"], d, config)
	transformed["tag"] = original["tag"]
	// flattenDialogflowCXFulfillmentTag(original["tag"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXFulfillmentMessages(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})

		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		var key = "text"

		if original[key] != nil {
			transformed = append(transformed, map[string]interface{}{
				key: flattenDialogflowCXFulfillmentMessagesText(original[key], d, config),
			})
		}
		key = "payload"
		if original[key] != nil {
			transformed = append(transformed, map[string]interface{}{
				key: flattenDialogflowCXFulfillmentMessagesPayload(original[key], d, config),
			})
		}
	}
	return transformed
}

func flattenDialogflowCXFulfillmentMessagesPayload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	originalText := original["text"]
	if originalText != nil {
		transformed["text"] = originalText
	}

	originalImage := original["image"]
	if originalImage != nil {
		transformed["image"] = originalImage
	}

	originalSimpleList := original["simple_list"]
	if originalSimpleList != nil {
		transformed["simple_list"] = originalSimpleList
	}
	originalQuickReplies := original["quick_replies"]
	if originalQuickReplies != nil {
		transformed["quick_replies"] = originalQuickReplies
	}
	originalListSuggestions := original["list_suggestions"]
	if originalListSuggestions != nil {
		transformed["list_suggestions"] = originalListSuggestions
	}

	return []interface{}{transformed}
}

func flattenDialogflowCXFulfillmentMessagesText(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	// transformed["text"] =
	// 	flattenDialogflowCXFulfillmentMessagesTextText(original["text"], d, config)
	transformed["text"] = original["text"]
	transformed["allow_playback_interruption"] = original["allowPlaybackInterruption"]
	return []interface{}{transformed}
}

// func flattenDialogflowCXFulfillmentMessagesText(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
// 	if v == nil {
// 		return nil
// 	}
// 	original := v.(map[string]interface{})
// 	if len(original) == 0 {
// 		return nil
// 	}
// 	transformed := make(map[string]interface{})
// 	originalText := original["text"]
// 	if originalText != nil {
// 		transformed["text"] = originalText
// 	}

// 	originalImage := original["image"]
// 	if originalImage != nil {
// 		transformed["image"] = originalImage
// 	}
// 	transformed["allow_playback_interruption"] = original["allowPlaybackInterruption"]
// 	return []interface{}{transformed}
// }

func expandDialogflowCXFulfillment(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMessages, err := expandDialogflowCXFulfillmentMessages(original["messages"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMessages); val.IsValid() && !isEmptyValue(val) {
		transformed["messages"] = transformedMessages
	}
	transformedSetParameterActions, err := expandDialogflowCXFulfillmentSetParameterActions(original["set_parameter_actions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSetParameterActions); val.IsValid() && !isEmptyValue(val) {
		transformed["setParameterActions"] = transformedSetParameterActions
	}
	transformedWebhook, err := expandDialogflowCXFulfillmentWebhook(original["webhook"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWebhook); val.IsValid() && !isEmptyValue(val) {
		transformed["webhook"] = transformedWebhook
	}

	transformedReturnPartialResponses, err := expandDialogflowCXFulfillmentReturnPartialResponses(original["return_partial_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReturnPartialResponses); val.IsValid() && !isEmptyValue(val) {
		transformed["returnPartialResponses"] = transformedReturnPartialResponses
	}

	transformedTag, err := expandDialogflowCXFulfillmentTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandDialogflowCXFulfillmentMessages(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		var key = "text"
		var originalValue = original[key].([]interface{})
		if len(originalValue) > 0 {
			transformedText, err := expandDialogflowCXFulfillmentMessagesText(original[key], d, config)
			if err != nil {
				return nil, err
			} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
				transformed[key] = transformedText
			}

			req = append(req, transformed)
		}

		key = "payload"
		originalValue = original[key].([]interface{})
		if len(originalValue) > 0 {
			transformedText, err := expandDialogflowCXFulfillmentMessagesPayload(original[key], d, config)
			if err != nil {
				return nil, err
			} else if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
				transformed[key] = transformedText
			}

			req = append(req, transformed)
		}

		// req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXFulfillmentMessagesPayload(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	expandDialogFlowCXValue(original, d, config, transformed, "text")
	expandDialogFlowCXValue(original, d, config, transformed, "image")
	expandDialogFlowCXValue(original, d, config, transformed, "simple_list")
	expandDialogFlowCXValue(original, d, config, transformed, "quick_replies")
	expandDialogFlowCXValue(original, d, config, transformed, "list_suggestions")

	return transformed, nil
}
func expandDialogflowCXFulfillmentSetParameterActions(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})
		var key = "parameter"
		var originalValue = original[key]
		if originalValue != nil {
			if val := reflect.ValueOf(originalValue); val.IsValid() && !isEmptyValue(val) {
				transformed[key] = originalValue
			}
		}

		key = "value"
		originalValue = original[key]
		if originalValue != nil {
			if val := reflect.ValueOf(originalValue); val.IsValid() && !isEmptyValue(val) {
				transformed[key] = originalValue
			}
		}
		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogFlowCXValue(original map[string]interface{}, d TerraformResourceData, config *transport_tpg.Config, transformed map[string]interface{}, key string) {
	transformedText := original[key]
	if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed[key] = transformedText
	}
}

func expandDialogflowCXFulfillmentMessagesText(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText := original["text"]
	if val := reflect.ValueOf(transformedText); val.IsValid() && !isEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedAllowPlaybackInterruption, err := expandDialogflowCXFulfillmentMessagesTextAllowPlaybackInterruption(original["allow_playback_interruption"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowPlaybackInterruption); val.IsValid() && !isEmptyValue(val) {
		transformed["allowPlaybackInterruption"] = transformedAllowPlaybackInterruption
	}

	return transformed, nil
}

func expandDialogflowCXFulfillmentMessagesTextAllowPlaybackInterruption(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFulfillmentWebhook(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFulfillmentReturnPartialResponses(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXFulfillmentTag(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func flattenDialogflowCXEventHandlers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":                flattenDialogflowCXEventHandlersName(original["name"], d, config),
			"event":               flattenDialogflowCXEventHandlersEvent(original["event"], d, config),
			"trigger_fulfillment": flattenDialogflowCXFulfillment(original["triggerFulfillment"], d, config),
			"target_page":         flattenDialogflowCXEventHandlersTargetPage(original["targetPage"], d, config),
			"target_flow":         flattenDialogflowCXEventHandlersTargetFlow(original["targetFlow"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowCXEventHandlersName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXEventHandlersEvent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXEventHandlersTargetPage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXEventHandlersTargetFlow(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDialogflowCXEventHandlers(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDialogflowCXEventHandlersName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedEvent, err := expandDialogflowCXEventHandlersEvent(original["event"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEvent); val.IsValid() && !isEmptyValue(val) {
			transformed["event"] = transformedEvent
		}

		transformedTriggerFulfillment, err := expandDialogflowCXFulfillment(original["trigger_fulfillment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTriggerFulfillment); val.IsValid() && !isEmptyValue(val) {
			transformed["triggerFulfillment"] = transformedTriggerFulfillment
		}

		transformedTargetPage, err := expandDialogflowCXEventHandlersTargetPage(original["target_page"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetPage); val.IsValid() && !isEmptyValue(val) {
			transformed["targetPage"] = transformedTargetPage
		}

		transformedTargetFlow, err := expandDialogflowCXEventHandlersTargetFlow(original["target_flow"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetFlow); val.IsValid() && !isEmptyValue(val) {
			transformed["targetFlow"] = transformedTargetFlow
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXEventHandlersName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEventHandlersEvent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEventHandlersTargetPage(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEventHandlersTargetFlow(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
