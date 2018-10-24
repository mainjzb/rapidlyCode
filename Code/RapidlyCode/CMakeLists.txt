cmake_minimum_required(VERSION 3.6.0)
Message("Into Build.cmake")

set(targetname "RapidlyCode")
set(OUT_PUT_DIRECTORY "../../../bin/")


set(CMAKE_CXX_STANDARD 11)

if(EXISTS "${TOOLS_CMAKE_DIR}/CustomConfig.cmake")
	include("${TOOLS_CMAKE_DIR}/CustomConfig.cmake")
else()
	set(QT_DIR "D:/qt/5.11.1/msvc2017_64/")
endif()


set(CMAKE_INCLUDE_CURRENT_DIR ON)
set(CMAKE_AUTOMOC ON)

set(Qt5_DIR "${QT_DIR}")

set(QT_DIR "${QT_DIR}" CACHE STRING "QT directory" FORCE)
set(QT5_DIR "${QT5_DIR}" CACHE STRING "QT directory" FORCE)
set(CMAKE_PREFIX_PATH ${cMAKE_PREFIX_PATH} "${QT5_DIR}")


find_package(Qt5Core REQUIRED)
find_package(Qt5Widgets REQUIRED)



file(GLOB Test_SRC
"*.h"
"*.cpp"
)
add_executable(RapidlyCode ${Test_SRC})

target_link_libraries(RapidlyCode Qt5::Widgets)
set_target_properties( ${targetname} PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${OUT_PUT_DIRECTORY} )
set_target_properties( ${targetname} PROPERTIES RUNTIME_OUTPUT_DIRECTORY_DEBUG ${OUT_PUT_DIRECTORY} )
set_target_properties( ${targetname} PROPERTIES RUNTIME_OUTPUT_DIRECTORY_RELEASE ${OUT_PUT_DIRECTORY} )


set_target_properties(${targetname} PROPERTIES OUTPUT_NAME_DEBUG "${targetname}_debug")
set_target_properties(${targetname} PROPERTIES OUTPUT_NAME_RELEASE "${targetname}_Release")