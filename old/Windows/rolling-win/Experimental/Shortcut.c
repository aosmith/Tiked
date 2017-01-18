#include <windows.h>
#include <stdio.h>

BOOL CreateShortcut(char *ptchExecutableFileName, char *ptchShortcutName)
{
    BOOL Res = TRUE;
    FILE *fp;
    
    if(NULL == (fp = fopen(ptchExecutableFileName, "r")))
    {
        printf("\nCan`t find executable file!\n");
        return FALSE;
    }
    fclose(fp);
    
    Res &= WritePrivateProfileString("InternetShortcut", 
               "URL", ptchExecutableFileName, ptchShortcutName);
    Res &= WritePrivateProfileString("InternetShortcut", 
               "IconIndex", "0", ptchShortcutName);
    Res &= WritePrivateProfileString("InternetShortcut", 
               "IconFile", ptchExecutableFileName, ptchShortcutName);
    
    return Res;
}


/* 
    For example: create a shortcut on desktop for e:\Test.exe 
    CreateShortcut("E:\\Test.exe", 
      "C:\\Documents and Settings\\Administrator\\Desktop\\Test.url");
*/