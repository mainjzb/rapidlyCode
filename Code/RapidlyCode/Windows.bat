cmake -E make_directory "./Solution" && cmake -E chdir "./Solution" cmake -D BUILD_SERVER=1 -G "Visual Studio 15 Win64" ../
pause