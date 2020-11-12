// +build cgo,windows,vulkan,!d3d,!d3d12,!opengl

package kinc

// #cgo CFLAGS: -I${SRCDIR}\..\Backends\Graphics5\Vulkan\Sources -I${SRCDIR}\..\Backends\Graphics4\G4onG5\Sources
// #cgo CPPFLAGS: -I${SRCDIR}\..\Backends\Graphics5\Vulkan\Sources -I${SRCDIR}\..\Backends\Graphics4\G4onG5\Sources
// #cgo CPPFLAGS: -I "${VULKAN_SDK}\Include"
// #cgo CPPFLAGS: -DKORE_G4=1 -DKORE_G5=1 -DKORE_G4ONG5
// #cgo CFLAGS: -DKORE_G4=1 -DKORE_G5=1 -DKORE_G4ONG5
// #cgo CPPFLAGS: -DKORE_VULKAN=1 -DVK_USE_PLATFORM_WIN32_KHR=1
// #cgo CFLAGS: -DKORE_VULKAN=1 -DVK_USE_PLATFORM_WIN32_KHR=1
import "C"

// todo:  get individual VKLayer lib
/**
		let libs = fs.readdirSync(path.join(process.env.VULKAN_SDK, 'Lib32'));
		for (const lib of libs) {
			if (lib.startsWith('VkLayer_')) {
				project.addLibFor('Win32', path.join(process.env.VULKAN_SDK, 'Lib32', lib.substr(0, lib.length - 4)));
			}
		}
		libs = fs.readdirSync(path.join(process.env.VULKAN_SDK, 'Lib'));
		for (const lib of libs) {
			if (lib.startsWith('VkLayer_')) {
				project.addLibFor('x64', path.join(process.env.VULKAN_SDK, 'Lib', lib.substr(0, lib.length - 4)));
			}
		}
		project.addIncludeDir(path.join(process.env.VULKAN_SDK, 'Include'));
**/
