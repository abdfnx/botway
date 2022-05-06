/**
 * Global configuration
 */
export const CONFIG = {
  /**
   * @type {string}
   */
  name: "botway",

  /**
   * @type {string}
   */
  path: "./bin",

  /**
   * @type {string}
   */
  url: "https://github.com/abdfnx/botway/releases/download/v{{version}}/{{bin_name}}_{{platform}}_v{{version}}_{{arch}}.zip",
};

/**
 * Mapping from Node's `process.arch` to Golang's `$GOARCH`
 */
export const ARCH_MAPPING = {
  ia32: "386",
  x64: "amd64",
  arm64: "arm64",
  arm: "arm",
};

/**
 * Mapping between Node's `process.platform` to Golang's
 */
export const PLATFORM_MAPPING = {
  darwin: "darwin",
  linux: "linux",
  win32: "windows",
  freebsd: "freebsd",
};
