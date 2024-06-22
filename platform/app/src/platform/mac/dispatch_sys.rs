/* automatically generated by rust-bindgen 0.65.1 */

pub const DISPATCH_TIME_NOW: u32 = 0;
pub const DISPATCH_QUEUE_PRIORITY_HIGH: u32 = 2;
pub type dispatch_function_t =
    ::std::option::Option<unsafe extern "C" fn(arg1: *mut ::std::os::raw::c_void)>;
pub type dispatch_time_t = u64;
extern "C" {
    pub fn dispatch_time(when: dispatch_time_t, delta: i64) -> dispatch_time_t;
}
#[repr(C)]
#[derive(Copy, Clone)]
pub union dispatch_object_t {
    pub _os_obj: *mut _os_object_s,
    pub _do: *mut dispatch_object_s,
    pub _dq: *mut dispatch_queue_s,
    pub _dqa: *mut dispatch_queue_attr_s,
    pub _dg: *mut dispatch_group_s,
    pub _ds: *mut dispatch_source_s,
    pub _dch: *mut dispatch_channel_s,
    pub _dm: *mut dispatch_mach_s,
    pub _dmsg: *mut dispatch_mach_msg_s,
    pub _dsema: *mut dispatch_semaphore_s,
    pub _ddata: *mut dispatch_data_s,
    pub _dchannel: *mut dispatch_io_s,
}
extern "C" {
    pub fn dispatch_set_context(object: dispatch_object_t, context: *mut ::std::os::raw::c_void);
}
extern "C" {
    pub fn dispatch_suspend(object: dispatch_object_t);
}
extern "C" {
    pub fn dispatch_resume(object: dispatch_object_t);
}
pub type dispatch_queue_t = *mut dispatch_queue_s;
pub type dispatch_queue_global_t = dispatch_queue_t;
extern "C" {
    pub fn dispatch_async_f(
        queue: dispatch_queue_t,
        context: *mut ::std::os::raw::c_void,
        work: dispatch_function_t,
    );
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_queue_s {
    pub _address: u8,
}
extern "C" {
    pub static mut _dispatch_main_q: dispatch_queue_s;
}
extern "C" {
    pub fn dispatch_get_global_queue(identifier: isize, flags: usize) -> dispatch_queue_global_t;
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_queue_attr_s {
    pub _address: u8,
}
extern "C" {
    pub fn dispatch_after_f(
        when: dispatch_time_t,
        queue: dispatch_queue_t,
        context: *mut ::std::os::raw::c_void,
        work: dispatch_function_t,
    );
}
pub type dispatch_source_t = *mut dispatch_source_s;
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_source_type_s {
    _unused: [u8; 0],
}
pub type dispatch_source_type_t = *const dispatch_source_type_s;
extern "C" {
    pub static _dispatch_source_type_data_add: dispatch_source_type_s;
}
extern "C" {
    pub fn dispatch_source_create(
        type_: dispatch_source_type_t,
        handle: usize,
        mask: usize,
        queue: dispatch_queue_t,
    ) -> dispatch_source_t;
}
extern "C" {
    pub fn dispatch_source_set_event_handler_f(
        source: dispatch_source_t,
        handler: dispatch_function_t,
    );
}
extern "C" {
    pub fn dispatch_source_cancel(source: dispatch_source_t);
}
extern "C" {
    pub fn dispatch_source_merge_data(source: dispatch_source_t, value: usize);
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_data_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct _os_object_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_object_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_group_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_source_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_channel_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_mach_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_mach_msg_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_semaphore_s {
    pub _address: u8,
}
#[repr(C)]
#[derive(Debug, Copy, Clone)]
pub struct dispatch_io_s {
    pub _address: u8,
}
