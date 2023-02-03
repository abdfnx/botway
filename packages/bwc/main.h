#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

#ifdef _WIN32
#define PATH_JOIN_SEPERATOR "\\"
#else
#define PATH_JOIN_SEPERATOR "/"
#endif

bool str_ends_with(const char *str, const char *end)
{
    int end_len;
    int str_len;

    if (NULL == str || NULL == end)
        return false;

    end_len = strlen(end);
    str_len = strlen(str);

    return str_len < end_len
               ? false
               : !strcmp(str + str_len - end_len, end);
}

bool str_starts_with(const char *str, const char *start)
{
    for (;; str++, start++)
        if (!*start)
            return true;
        else if (*str != *start)
            return false;
}

char *path_join(const char *dir, const char *file)
{
    int size = strlen(dir) + strlen(file) + 2;

    char *buf = malloc(size * sizeof(char));

    if (NULL == buf)
        return NULL;

    strcpy(buf, dir);

    if (!str_ends_with(dir, PATH_JOIN_SEPERATOR))
    {
        strcat(buf, PATH_JOIN_SEPERATOR);
    }

    if (str_starts_with(file, PATH_JOIN_SEPERATOR))
    {
        char *filecopy = strdup(file);

        if (NULL == filecopy)
        {
            free(buf);

            return NULL;
        }

        strcat(buf, ++filecopy);

        free(--filecopy);
    }
    else
    {
        strcat(buf, file);
    }

    return buf;
}
