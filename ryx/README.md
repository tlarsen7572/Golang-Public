## Refactoryx

Refactoryx is an application for interacting with Alteryx workflows on a project/library basis.  It is meant to complement Designer with enhanced macro and folder operations, tasks that are difficult and prone to error today.  Some examples include:
- Convert all of your macros in all of your workflows to relative or absolute file paths
- Move macros and folders of macros to new locations and update the macro references in all workflows that use them
- Select tools in a workflow and extract them to a macro

**Refactoryx is NOT complete and should not be used in production!  It is under active development; there are no alpha, beta, or production releases available at this time.**

Multiple repositories comprise the entire Refactoryx application.  A unified project view of all outstanding enhancements and bugs is provided by the [ryx GitHub project](https://github.com/users/tlarsen7572/projects/1).

## ryx

ryx is the back-end server and main entry point for the Refactoryx application.  It is built using Go and implements the majority of the business logic while also serving the [web-based front-end GUI](https://github.com/tlarsen7572/ryx_gui).  A [small .NET sub-application in another repository](https://github.com/tlarsen7572/IconLoader) is used to load tool connection data and icons from Alteryx's DLL files.

ryx uses a JSON configuration file to determine several important runtime parameters:
- InstallPath: This is the installation path to Alteryx.
- ProgramDataPath: This is the path to Alteryx's ProgramData folder where global Alteryx configuration settings are stored.
- UserFolders: A list of paths to each user's home directory on Windows.  Any users not included in this setting will not see their user-specific custom tools render properly in Refactoryx.
- HttpPort: The port on which to serve the front-end GUI.
- BrowseFolderRoots: A list of folders on the local machine.  This setting limits users to selecting projects inside these folders.  A typical practice might be to create a folder at C:\AlteryxProjects which will contain all of the Alteryx project folders.  Setting BrowseFolderRoots will limit users to selecting folders inside C:\AlteryxProjects and will prevent them from accessing other folders such as C:\Users and C:\Windows.  This setting is required.
- LogPath: The path to the ryx log file.  The log file generally stay empty unless a critical error has caused the application to shut down.  Any existing log is deleted when Refactoryx is started.

