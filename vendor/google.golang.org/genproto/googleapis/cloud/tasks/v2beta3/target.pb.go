// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2beta3/target.proto

package tasks // import "google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The HTTP method used to execute the task.
type HttpMethod int32

const (
	// HTTP method unspecified
	HttpMethod_HTTP_METHOD_UNSPECIFIED HttpMethod = 0
	// HTTP POST
	HttpMethod_POST HttpMethod = 1
	// HTTP GET
	HttpMethod_GET HttpMethod = 2
	// HTTP HEAD
	HttpMethod_HEAD HttpMethod = 3
	// HTTP PUT
	HttpMethod_PUT HttpMethod = 4
	// HTTP DELETE
	HttpMethod_DELETE HttpMethod = 5
)

var HttpMethod_name = map[int32]string{
	0: "HTTP_METHOD_UNSPECIFIED",
	1: "POST",
	2: "GET",
	3: "HEAD",
	4: "PUT",
	5: "DELETE",
}

var HttpMethod_value = map[string]int32{
	"HTTP_METHOD_UNSPECIFIED": 0,
	"POST":                    1,
	"GET":                     2,
	"HEAD":                    3,
	"PUT":                     4,
	"DELETE":                  5,
}

func (x HttpMethod) String() string {
	return proto.EnumName(HttpMethod_name, int32(x))
}

func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{0}
}

// App Engine HTTP queue.
//
// The task will be delivered to the App Engine application hostname
// specified by its [AppEngineHttpQueue][google.cloud.tasks.v2beta3.AppEngineHttpQueue] and [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest].
// The documentation for [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest] explains how the
// task's host URL is constructed.
//
// Using [AppEngineHttpQueue][google.cloud.tasks.v2beta3.AppEngineHttpQueue] requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
type AppEngineHttpQueue struct {
	// Overrides for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
	//
	// If set, `app_engine_routing_override` is used for all tasks in
	// the queue, no matter what the setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride *AppEngineRouting `protobuf:"bytes,1,opt,name=app_engine_routing_override,json=appEngineRoutingOverride,proto3" json:"app_engine_routing_override,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *AppEngineHttpQueue) Reset()         { *m = AppEngineHttpQueue{} }
func (m *AppEngineHttpQueue) String() string { return proto.CompactTextString(m) }
func (*AppEngineHttpQueue) ProtoMessage()    {}
func (*AppEngineHttpQueue) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{0}
}
func (m *AppEngineHttpQueue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineHttpQueue.Unmarshal(m, b)
}
func (m *AppEngineHttpQueue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineHttpQueue.Marshal(b, m, deterministic)
}
func (m *AppEngineHttpQueue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineHttpQueue.Merge(m, src)
}
func (m *AppEngineHttpQueue) XXX_Size() int {
	return xxx_messageInfo_AppEngineHttpQueue.Size(m)
}
func (m *AppEngineHttpQueue) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineHttpQueue.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineHttpQueue proto.InternalMessageInfo

func (m *AppEngineHttpQueue) GetAppEngineRoutingOverride() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRoutingOverride
	}
	return nil
}

// App Engine HTTP request.
//
// The message defines the HTTP request that is sent to an App Engine app when
// the task is dispatched.
//
// This proto can only be used for tasks in a queue which has
// [app_engine_http_queue][google.cloud.tasks.v2beta3.Queue.app_engine_http_queue] set.
//
// Using [AppEngineHttpRequest][google.cloud.tasks.v2beta3.AppEngineHttpRequest] requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
//
// The task will be delivered to the App Engine app which belongs to the same
// project as the queue. For more information, see
// [How Requests are Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
// and how routing is affected by
// [dispatch files](https://cloud.google.com/appengine/docs/python/config/dispatchref).
//
// The [AppEngineRouting][google.cloud.tasks.v2beta3.AppEngineRouting] used to construct the URL that the task is
// delivered to can be set at the queue-level or task-level:
//
// * If set,
//    [app_engine_routing_override][google.cloud.tasks.v2beta3.AppEngineHttpQueue.app_engine_routing_override]
//    is used for all tasks in the queue, no matter what the setting
//    is for the
//    [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
//
//
// The `url` that the task will be sent to is:
//
// * `url =` [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] `+`
//   [relative_uri][google.cloud.tasks.v2beta3.AppEngineHttpRequest.relative_uri]
//
// The task attempt has succeeded if the app's request handler returns
// an HTTP response code in the range [`200` - `299`]. `503` is
// considered an App Engine system error instead of an application
// error. Requests returning error `503` will be retried regardless of
// retry configuration and not counted against retry counts.
// Any other response code or a failure to receive a response before the
// deadline is a failed attempt.
type AppEngineHttpRequest struct {
	// The HTTP method to use for the request. The default is POST.
	//
	// The app's request handler for the task's target URL must be able to handle
	// HTTP requests with this http_method, otherwise the task attempt will fail
	// with error code 405 (Method Not Allowed). See
	// [Writing a push task request handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler)
	// and the documentation for the request handlers in the language your app is
	// written in e.g.
	// [Python Request Handler](https://cloud.google.com/appengine/docs/python/tools/webapp/requesthandlerclass).
	HttpMethod HttpMethod `protobuf:"varint,1,opt,name=http_method,json=httpMethod,proto3,enum=google.cloud.tasks.v2beta3.HttpMethod" json:"http_method,omitempty"`
	// Task-level setting for App Engine routing.
	//
	// If set,
	// [app_engine_routing_override][google.cloud.tasks.v2beta3.AppEngineHttpQueue.app_engine_routing_override]
	// is used for all tasks in the queue, no matter what the setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta3.AppEngineHttpRequest.app_engine_routing].
	AppEngineRouting *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing,json=appEngineRouting,proto3" json:"app_engine_routing,omitempty"`
	// The relative URI.
	//
	// The relative URI must begin with "/" and must be a valid HTTP relative URI.
	// It can contain a path and query string arguments.
	// If the relative URI is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters.
	RelativeUri string `protobuf:"bytes,3,opt,name=relative_uri,json=relativeUri,proto3" json:"relative_uri,omitempty"`
	// HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the
	// [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	// Repeated headers are not supported but a header value can contain commas.
	//
	// Cloud Tasks sets some headers to default values:
	//
	// * `User-Agent`: By default, this header is
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"`.
	//   This header can be modified, but Cloud Tasks will append
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"` to the
	//   modified `User-Agent`.
	//
	// If the task has a [body][google.cloud.tasks.v2beta3.AppEngineHttpRequest.body], Cloud
	// Tasks sets the following headers:
	//
	// * `Content-Type`: By default, the `Content-Type` header is set to
	//   `"application/octet-stream"`. The default can be overridden by explicitly
	//   setting `Content-Type` to a particular media type when the
	//   [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//   For example, `Content-Type` can be set to `"application/json"`.
	// * `Content-Length`: This is computed by Cloud Tasks. This value is
	//   output only.   It cannot be changed.
	//
	// The headers below cannot be set or overridden:
	//
	// * `Host`
	// * `X-Google-*`
	// * `X-AppEngine-*`
	//
	// In addition, Cloud Tasks sets some headers when the task is dispatched,
	// such as headers containing information about the task; see
	// [request headers](https://cloud.google.com/appengine/docs/python/taskqueue/push/creating-handlers#reading_request_headers).
	// These headers are set only when the task is dispatched, so they are not
	// visible when the task is returned in a Cloud Tasks response.
	//
	// Although there is no specific limit for the maximum number of headers or
	// the size, there is a limit on the maximum size of the [Task][google.cloud.tasks.v2beta3.Task]. For more
	// information, see the [CreateTask][google.cloud.tasks.v2beta3.CloudTasks.CreateTask] documentation.
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// HTTP request body.
	//
	// A request body is allowed only if the HTTP method is POST or PUT. It is
	// an error to set a body on a task with an incompatible [HttpMethod][google.cloud.tasks.v2beta3.HttpMethod].
	Body                 []byte   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineHttpRequest) Reset()         { *m = AppEngineHttpRequest{} }
func (m *AppEngineHttpRequest) String() string { return proto.CompactTextString(m) }
func (*AppEngineHttpRequest) ProtoMessage()    {}
func (*AppEngineHttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{1}
}
func (m *AppEngineHttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineHttpRequest.Unmarshal(m, b)
}
func (m *AppEngineHttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineHttpRequest.Marshal(b, m, deterministic)
}
func (m *AppEngineHttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineHttpRequest.Merge(m, src)
}
func (m *AppEngineHttpRequest) XXX_Size() int {
	return xxx_messageInfo_AppEngineHttpRequest.Size(m)
}
func (m *AppEngineHttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineHttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineHttpRequest proto.InternalMessageInfo

func (m *AppEngineHttpRequest) GetHttpMethod() HttpMethod {
	if m != nil {
		return m.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (m *AppEngineHttpRequest) GetAppEngineRouting() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRouting
	}
	return nil
}

func (m *AppEngineHttpRequest) GetRelativeUri() string {
	if m != nil {
		return m.RelativeUri
	}
	return ""
}

func (m *AppEngineHttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AppEngineHttpRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// App Engine Routing.
//
// Specifies the target URI. Since this target type dispatches tasks to secure
// app handlers, unsecure app handlers, and URIs restricted with
// [`login: admin`](https://cloud.google.com/appengine/docs/standard/python/config/appref)
// the protocol (for example, HTTP or HTTPS) cannot be explictly specified.
// Task dispatches do not follow redirects and cannot target URI paths
// restricted with
// [`login: required`](https://cloud.google.com/appengine/docs/standard/python/config/appref)
// because tasks are not run as any user.
//
// For more information about services, versions, and instances see
// [An Overview of App Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
type AppEngineRouting struct {
	// App service.
	//
	// By default, the task is sent to the service which is the default
	// service when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable
	// into [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance]. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable, then
	// [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance] are the empty string.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// App version.
	//
	// By default, the task is sent to the version which is the default
	// version when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable
	// into [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance]. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta3.AppEngineRouting.host] is not parsable, then
	// [service][google.cloud.tasks.v2beta3.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta3.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance] are the empty string.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// App instance.
	//
	// By default, the task is sent to an instance which is available when
	// the task is attempted.
	//
	// Requests can only be sent to a specific instance if
	// [manual scaling is used in App Engine Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).
	// App Engine Flex does not support instances. For more information, see
	// [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
	// and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
	Instance string `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	// Output only. The host that the task is sent to.
	//
	// The host is constructed from the domain name of the app associated with
	// the queue's project ID (for example <app-id>.appspot.com), and the
	// [service][google.cloud.tasks.v2beta3.AppEngineRouting.service], [version][google.cloud.tasks.v2beta3.AppEngineRouting.version],
	// and [instance][google.cloud.tasks.v2beta3.AppEngineRouting.instance]. Tasks which were created using
	// the App Engine SDK might have a custom domain name.
	//
	// For more information, see
	// [How Requests are Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppEngineRouting) Reset()         { *m = AppEngineRouting{} }
func (m *AppEngineRouting) String() string { return proto.CompactTextString(m) }
func (*AppEngineRouting) ProtoMessage()    {}
func (*AppEngineRouting) Descriptor() ([]byte, []int) {
	return fileDescriptor_595de6119aed6d9e, []int{2}
}
func (m *AppEngineRouting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppEngineRouting.Unmarshal(m, b)
}
func (m *AppEngineRouting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppEngineRouting.Marshal(b, m, deterministic)
}
func (m *AppEngineRouting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppEngineRouting.Merge(m, src)
}
func (m *AppEngineRouting) XXX_Size() int {
	return xxx_messageInfo_AppEngineRouting.Size(m)
}
func (m *AppEngineRouting) XXX_DiscardUnknown() {
	xxx_messageInfo_AppEngineRouting.DiscardUnknown(m)
}

var xxx_messageInfo_AppEngineRouting proto.InternalMessageInfo

func (m *AppEngineRouting) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AppEngineRouting) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AppEngineRouting) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *AppEngineRouting) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*AppEngineHttpQueue)(nil), "google.cloud.tasks.v2beta3.AppEngineHttpQueue")
	proto.RegisterType((*AppEngineHttpRequest)(nil), "google.cloud.tasks.v2beta3.AppEngineHttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.tasks.v2beta3.AppEngineHttpRequest.HeadersEntry")
	proto.RegisterType((*AppEngineRouting)(nil), "google.cloud.tasks.v2beta3.AppEngineRouting")
	proto.RegisterEnum("google.cloud.tasks.v2beta3.HttpMethod", HttpMethod_name, HttpMethod_value)
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2beta3/target.proto", fileDescriptor_595de6119aed6d9e)
}

var fileDescriptor_595de6119aed6d9e = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x51, 0x6b, 0xd3, 0x50,
	0x14, 0x36, 0x4d, 0xbb, 0xae, 0xa7, 0x45, 0xc2, 0x65, 0x60, 0xe8, 0x44, 0x6a, 0x1f, 0xb4, 0x88,
	0x24, 0xd0, 0xbd, 0xc8, 0x44, 0x64, 0xb3, 0x71, 0x1d, 0x38, 0x1b, 0xb3, 0x14, 0x61, 0x3e, 0x84,
	0xdb, 0xf6, 0x90, 0x5e, 0xda, 0xdd, 0x1b, 0x6f, 0x6e, 0x02, 0x7d, 0xf4, 0xc9, 0xbf, 0x2d, 0xb9,
	0x49, 0xeb, 0xac, 0x3a, 0xc4, 0xb7, 0xf3, 0x7d, 0xf9, 0xce, 0x77, 0xee, 0xf9, 0xc2, 0x81, 0xe7,
	0xb1, 0x10, 0xf1, 0x1a, 0xdd, 0xf9, 0x5a, 0x64, 0x0b, 0x57, 0xd1, 0x74, 0x95, 0xba, 0xf9, 0x70,
	0x86, 0x8a, 0x9e, 0xb8, 0x8a, 0xca, 0x18, 0x95, 0x93, 0x48, 0xa1, 0x04, 0xe9, 0x96, 0x42, 0x47,
	0x0b, 0x1d, 0x2d, 0x74, 0x2a, 0x61, 0xf7, 0x71, 0x65, 0x42, 0x13, 0xe6, 0x52, 0xce, 0x85, 0xa2,
	0x8a, 0x09, 0x9e, 0x96, 0x9d, 0xfd, 0x6f, 0x06, 0x90, 0xb3, 0x24, 0xf1, 0x78, 0xcc, 0x38, 0x8e,
	0x95, 0x4a, 0x3e, 0x65, 0x98, 0x21, 0x59, 0xc1, 0x31, 0x4d, 0x92, 0x08, 0x35, 0x1d, 0x49, 0x91,
	0x29, 0xc6, 0xe3, 0x48, 0xe4, 0x28, 0x25, 0x5b, 0xa0, 0x6d, 0xf4, 0x8c, 0x41, 0x7b, 0xf8, 0xd2,
	0xf9, 0xfb, 0x58, 0x67, 0x67, 0x1a, 0x94, 0xcd, 0x81, 0x4d, 0xf7, 0x98, 0x49, 0xe5, 0xd6, 0xff,
	0x6e, 0xc2, 0xd1, 0x2f, 0x6f, 0x08, 0xf0, 0x6b, 0x86, 0xa9, 0x22, 0x17, 0xd0, 0x5e, 0x2a, 0x95,
	0x44, 0xb7, 0xa8, 0x96, 0x62, 0xa1, 0xa7, 0x3e, 0x1c, 0x3e, 0xbb, 0x6f, 0x6a, 0xd1, 0x7d, 0xa5,
	0xd5, 0x01, 0x2c, 0x77, 0x35, 0xb9, 0x01, 0xf2, 0xfb, 0x3a, 0x76, 0xed, 0x3f, 0xb6, 0xb0, 0xf6,
	0xb7, 0x20, 0x4f, 0xa1, 0x23, 0x71, 0x4d, 0x15, 0xcb, 0x31, 0xca, 0x24, 0xb3, 0xcd, 0x9e, 0x31,
	0x68, 0x05, 0xed, 0x2d, 0x37, 0x95, 0x8c, 0x7c, 0x86, 0xe6, 0x12, 0xe9, 0x02, 0x65, 0x6a, 0xd7,
	0x7b, 0xe6, 0xa0, 0x3d, 0x7c, 0xf3, 0x4f, 0x33, 0xef, 0x44, 0xe1, 0x8c, 0xcb, 0x7e, 0x8f, 0x2b,
	0xb9, 0x09, 0xb6, 0x6e, 0x84, 0x40, 0x7d, 0x26, 0x16, 0x1b, 0xbb, 0xd1, 0x33, 0x06, 0x9d, 0x40,
	0xd7, 0xdd, 0x53, 0xe8, 0xdc, 0x15, 0x13, 0x0b, 0xcc, 0x15, 0x6e, 0x74, 0x78, 0xad, 0xa0, 0x28,
	0xc9, 0x11, 0x34, 0x72, 0xba, 0xce, 0x50, 0x07, 0xd0, 0x0a, 0x4a, 0x70, 0x5a, 0x7b, 0x65, 0xf4,
	0x73, 0xb0, 0xf6, 0x37, 0x26, 0x36, 0x34, 0x53, 0x94, 0x39, 0x9b, 0x63, 0xe5, 0xb1, 0x85, 0xc5,
	0x97, 0x1c, 0x65, 0xca, 0x04, 0xaf, 0x9c, 0xb6, 0x90, 0x74, 0xe1, 0x90, 0xf1, 0x54, 0x51, 0x3e,
	0xc7, 0x2a, 0x8f, 0x1d, 0x2e, 0xde, 0xbc, 0x14, 0xa9, 0xb2, 0xeb, 0x9a, 0xd7, 0xf5, 0x8b, 0x2f,
	0x00, 0x3f, 0xff, 0x1c, 0x39, 0x86, 0x47, 0xe3, 0x30, 0xf4, 0xa3, 0x2b, 0x2f, 0x1c, 0x4f, 0x46,
	0xd1, 0xf4, 0xe3, 0xb5, 0xef, 0xbd, 0xbb, 0x7c, 0x7f, 0xe9, 0x8d, 0xac, 0x07, 0xe4, 0x10, 0xea,
	0xfe, 0xe4, 0x3a, 0xb4, 0x0c, 0xd2, 0x04, 0xf3, 0xc2, 0x0b, 0xad, 0x5a, 0x41, 0x8d, 0xbd, 0xb3,
	0x91, 0x65, 0x16, 0x94, 0x3f, 0x0d, 0xad, 0x3a, 0x01, 0x38, 0x18, 0x79, 0x1f, 0xbc, 0xd0, 0xb3,
	0x1a, 0xe7, 0x09, 0x3c, 0x99, 0x8b, 0xdb, 0x7b, 0x12, 0x3f, 0x6f, 0x87, 0xfa, 0x98, 0xfc, 0xe2,
	0x22, 0x7c, 0xe3, 0xe6, 0x6d, 0x25, 0x8d, 0xc5, 0x9a, 0xf2, 0xd8, 0x11, 0x32, 0x76, 0x63, 0xe4,
	0xfa, 0x5e, 0xdc, 0xf2, 0x13, 0x4d, 0x58, 0xfa, 0xa7, 0xab, 0x7c, 0xad, 0xd1, 0xec, 0x40, 0x6b,
	0x4f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x26, 0x4d, 0x45, 0xc0, 0x03, 0x00, 0x00,
}
