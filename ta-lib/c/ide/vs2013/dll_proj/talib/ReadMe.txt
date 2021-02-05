========================================================================
    动态链接库：ta-lib 项目概述
========================================================================
[2021-02-05] added by Jim4work
Q:How to fix Link Errors? (怎样修复连接错误[LNK2001]?)
  Error LNK2001: unresolved external symbol __TA_MACD ... (连接错误LNK2001:无法解析的外部符号: __TA_MACD 等等...)
A: You need to configure your project in Visual Studio. (您需要在 Visual Studio 中修改项目的配置属性)
  In your project properties, under Linker -> General, add '..\..\..\..\lib;' to Additional Library Directories. 
  (打开项目的配置属性, 找到'链接器'->'常规', 在'附加库目录'中添加'..\..\..\..\lib;')
  Under Linker -> Input, add 'ta_libc_cdr.lib;ta_libc_cmr.lib;'(and/or other ta_libc_**.libs) to Additional Dependencies. 
  (再找到'链接器'->'输入', 在'附加依赖项'中添加'ta_libc_cdr.lib;ta_libc_cmr.lib;'或其它几个库文件)
  This tells Visual Studio the name and location of your thirdparty libraries, so that it can link it into your main application.
  (这是告诉 Visual Studio 第三方库文件的名字和路径, 这样才能让编译器正确地将库链接到主程序中)    

应用程序向导已为您创建了此 ta-lib DLL。

本文件概要介绍组成 ta-lib 应用程序的每个文件的内容。


ta-lib.vcxproj
    这是使用应用程序向导生成的 VC++ 项目的主项目文件，其中包含生成该文件的 Visual C++ 的版本信息，以及有关使用应用程序向导选择的平台、配置和项目功能的信息。

ta-lib.vcxproj.filters
    这是使用“应用程序向导”生成的 VC++ 项目筛选器文件。它包含有关项目文件与筛选器之间的关联信息。在 IDE 中，通过这种关联，在特定节点下以分组形式显示具有相似扩展名的文件。例如，“.cpp”文件与“源文件”筛选器关联。

ta-lib.cpp
    这是主 DLL 源文件。

	此 DLL 在创建时不导出任何符号。因此，生成时不会产生 .lib 文件。如果希望此项目成为其他某个项目的项目依赖项，则需要添加代码以从 DLL 导出某些符号，以便产生一个导出库，或者，也可以在项目“属性页”对话框中的“链接器”文件夹中，将“常规”属性页上的“忽略输入库”属性设置为“是”。

/////////////////////////////////////////////////////////////////////////////
其他标准文件:

StdAfx.h, StdAfx.cpp
    这些文件用于生成名为 ta-lib.pch 的预编译头 (PCH) 文件和名为 StdAfx.obj 的预编译类型文件。

/////////////////////////////////////////////////////////////////////////////
其他注释:

应用程序向导使用“TODO:”注释来指示应添加或自定义的源代码部分。

/////////////////////////////////////////////////////////////////////////////
