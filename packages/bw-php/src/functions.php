<?php
namespace Botway;

function pathJoin($path1, $path2) {
    $paths = func_get_args();
    $last_key = func_num_args() - 1;

    array_walk($paths, function (&$val, $key) use ($last_key) {
        switch ($key) {
            case 0:
                $val = rtrim($val, "/ ");
                break;

            case $last_key:
                $val = ltrim($val, "/ ");
                break;

            default:
                $val = trim($val, "/ ");

                break;
        }
    });

    $first = array_shift($paths);
    $last = array_pop($paths);
    $paths = array_filter($paths);

    array_unshift($paths, $first);

    $paths[] = $last;

    return implode("/", $paths);
}

function homeDir() {
    if(isset($_SERVER["HOME"])) {
        $result = $_SERVER["HOME"];
    } else {
        $result = getenv("HOME");
    }

    if(empty($result) && function_exists("exec")) {
        if(strncasecmp(PHP_OS, "WIN", 3) === 0) {
            $result = exec("echo %userprofile%");
        } else {
            $result = exec("echo ~");
        }
    }

    return $result;
}
