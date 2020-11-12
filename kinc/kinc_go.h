#pragma once

#include "../Sources/kinc/pch.h"
#include "../Sources/kinc/window.h"
#include "../Sources/kinc/system.h"
#include "../Sources/kinc/image.h"
#include "../Sources/kinc/display.h"

#ifdef __cplusplus
extern "C" {
#endif

static bool toBool(int i)
{
 	if(i == 1) return true;
		return false;
}

static int fromBool(bool b)
{
 	if(b) return 1;
		return 0;
}

void resize_callback_proxy(int x, int y, void *data);
static void _resize_callback(int window, void *data) {
  kinc_window_set_resize_callback(window, resize_callback_proxy, data);
}

void ppi_changed_callback_proxy(int ppi, void *data);
static void _ppi_changed_callback(int window, void *data) {
  kinc_window_set_ppi_changed_callback(window, ppi_changed_callback_proxy, data);
}

void update_cb_proxy();
static void set_update_cb(){
	 kinc_set_update_callback(update_cb_proxy);
}

void foreground_cb_proxy();
static void set_foreground_cb(){
	 kinc_set_foreground_callback(foreground_cb_proxy);
}

void resume_cb_proxy();
static void set_resume_cb(){
	 kinc_set_resume_callback(resume_cb_proxy);
}

void pause_cb_proxy();
static void set_pause_cb(){
	 kinc_set_pause_callback(pause_cb_proxy);
}

void background_cb_proxy();
static void set_background_cb(){
	 kinc_set_background_callback(background_cb_proxy);
}

void shutdown_cb_proxy();
static void set_shutdown_cb(){
	 kinc_set_shutdown_callback(shutdown_cb_proxy);
}

void dropfiles_cb_proxy(wchar_t *data);
static void set_dropfiles_cb(){
	 kinc_set_drop_files_callback(dropfiles_cb_proxy);
}

char* cut_cb_proxy();
static void set_cut_cb(){
	 kinc_set_cut_callback(cut_cb_proxy);
}

char* copy_cb_proxy();
static void set_copy_cb(){
	 kinc_set_copy_callback(copy_cb_proxy);
}

void paste_cb_proxy(char *data);
static void set_paste_cb(){
	 kinc_set_paste_callback(paste_cb_proxy);
}

void login_cb_proxy();
static void set_login_cb(){
	 kinc_set_login_callback(login_cb_proxy);
}

void logout_cb_proxy();
static void set_logout_cb(){
	 kinc_set_logout_callback(logout_cb_proxy);
}

static int __kinc_count_displays(){
	return kinc_count_displays();
}
static int __kinc_primary_display(){
	return kinc_primary_display();
}
static int __kinc_display_available(int display_index){
		return fromBool(kinc_display_available(display_index));
}
static const char *__kinc_display_name(int display_index){
		return kinc_display_name(display_index);
}
static kinc_display_mode_t __kinc_display_current_mode(int display_index){
		return kinc_display_current_mode(display_index);
}
static int __kinc_display_count_available_modes(int display_index){
		return kinc_display_count_available_modes(display_index);
}
static kinc_display_mode_t __kinc_display_available_mode(int display_index, int mode_index){
		return kinc_display_available_mode(display_index, mode_index);
}

#ifdef __cplusplus
}
#endif