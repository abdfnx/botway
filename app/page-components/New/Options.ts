export const platforms = [
  {
    name: "Discord",
    slug: "discord",
  },
  {
    name: "Slack",
    slug: "slack",
  },
  {
    name: "Telegram",
    slug: "telegram",
  },
  {
    name: "Twitch",
    slug: "twitch",
  },
];

export const visibilityOptions = [
  {
    typeName: "Public",
    type: "public",
  },
  {
    typeName: "Private",
    type: "private",
  },
];

export const langs = (platform: any) => {
  const bwLangs = [
    {
      name: "Node.js",
      slug: "nodejs",
    },
    {
      name: "C++",
      slug: "cpp",
    },
    {
      name: "Crystal",
      slug: "crystal",
    },
    {
      name: "C#",
      slug: "csharp",
    },
    {
      name: "Dart",
      slug: "dart",
    },
    {
      name: "Deno",
      slug: "deno",
    },
    {
      name: "Go",
      slug: "go",
    },
    {
      name: "Java",
      slug: "java",
    },
    {
      name: "Kotlin",
      slug: "kotlin",
    },
    {
      name: "Nim",
      slug: "nim",
    },
    {
      name: "PHP",
      slug: "php",
    },
    {
      name: "Python",
      slug: "python",
    },
    {
      name: "Ruby",
      slug: "ruby",
    },
    {
      name: "Rust",
      slug: "rust",
    },
    {
      name: "Swift",
      slug: "swift",
    },
    {
      name: "TypeScript",
      slug: "typescript",
    },
  ];

  const discordLangs = [
    {
      name: "Node.js",
      slug: "nodejs",
    },
    {
      name: "C",
      slug: "c",
    },
    {
      name: "C++",
      slug: "cpp",
    },
    {
      name: "Crystal",
      slug: "crystal",
    },
    {
      name: "C#",
      slug: "csharp",
    },
    {
      name: "Dart",
      slug: "dart",
    },
    {
      name: "Deno",
      slug: "deno",
    },
    {
      name: "Go",
      slug: "go",
    },
    {
      name: "Java",
      slug: "java",
    },
    {
      name: "Kotlin",
      slug: "kotlin",
    },
    {
      name: "Nim",
      slug: "nim",
    },
    {
      name: "PHP",
      slug: "php",
    },
    {
      name: "Python",
      slug: "python",
    },
    {
      name: "Ruby",
      slug: "ruby",
    },
    {
      name: "Rust",
      slug: "rust",
    },
    {
      name: "TypeScript",
      slug: "typescript",
    },
  ];

  const telegramLangs = [
    {
      name: "Node.js",
      slug: "nodejs",
    },
    {
      name: "C++",
      slug: "cpp",
    },
    {
      name: "Crystal",
      slug: "crystal",
    },
    {
      name: "C#",
      slug: "csharp",
    },
    {
      name: "Dart",
      slug: "dart",
    },
    {
      name: "Deno",
      slug: "deno",
    },
    {
      name: "Go",
      slug: "go",
    },
    {
      name: "Java",
      slug: "java",
    },
    {
      name: "Kotlin",
      slug: "kotlin",
    },
    {
      name: "Nim",
      slug: "nim",
    },
    {
      name: "PHP",
      slug: "php",
    },
    {
      name: "Python",
      slug: "python",
    },
    {
      name: "Ruby",
      slug: "ruby",
    },
    {
      name: "Rust",
      slug: "rust",
    },
    {
      name: "Swift",
      slug: "swift",
    },
    {
      name: "TypeScript",
      slug: "typescript",
    },
  ];

  const slackLangs = [
    {
      name: "Node.js",
      slug: "nodejs",
    },
    {
      name: "Python",
      slug: "python",
    },
    {
      name: "TypeScript",
      slug: "typescript",
    },
  ];

  const twitchLangs = [
    {
      name: "Node.js",
      slug: "nodejs",
    },

    {
      name: "Deno",
      slug: "deno",
    },
    {
      name: "Go",
      slug: "go",
    },
    {
      name: "Java",
      slug: "java",
    },
    {
      name: "Python",
      slug: "python",
    },
    {
      name: "TypeScript",
      slug: "typescript",
    },
  ];

  switch (platform) {
    case "Discord":
      return discordLangs;

    case "Telegram":
      return telegramLangs;

    case "Slack":
      return slackLangs;

    case "Twitch":
      return twitchLangs;

    default:
      return bwLangs;
  }
};

export const hostServices = [
  {
    name: "Railway",
    logo: "railway.svg",
  },
  {
    name: "Render",
    logo: "render.png",
  },
];

export const packageManagers = (lang: any) => {
  const nodePMs = [
    {
      name: "npm",
      logo: "npm.svg",
    },
    {
      name: "pnpm",
      logo: "pnpm.svg",
    },
    {
      name: "yarn",
      logo: "yarn.svg",
    },
  ];

  const cPM = [
    {
      name: "default",
      logo: "c.svg",
    },
  ];

  const cppPM = [
    {
      name: "cmake",
      logo: "cmake.svg",
    },
  ];

  const crytsalPM = [
    {
      name: "shards",
      logo: "crystal.svg",
    },
  ];

  const csharpPM = [
    {
      name: "dotnet",
      logo: "dotnet.svg",
    },
  ];

  const dartPM = [
    {
      name: "pub",
      logo: "dart.svg",
    },
  ];

  const denoPM = [
    {
      name: "deno package manager",
      logo: "deno.svg",
    },
  ];

  const goPM = [
    {
      name: "go package manager",
      logo: "go.svg",
    },
  ];

  const gradlePM = [
    {
      name: "gradle",
      logo: "gradle.svg",
    },
  ];

  const nimPM = [
    {
      name: "nimble",
      logo: "nimble.svg",
    },
  ];

  const phpPM = [
    {
      name: "composer",
      logo: "composer.svg",
    },
  ];

  const pythonPMs = [
    {
      name: "pip",
      logo: "pip.svg",
    },
    {
      name: "pipenv",
      logo: "pipenv.png",
    },
    {
      name: "poetry",
      logo: "poetry.svg",
    },
  ];

  const rubyPM = [
    {
      name: "bundler",
      logo: "bundler.png",
    },
  ];

  const rustPMs = [
    {
      name: "cargo",
      logo: "cargo.png",
    },
    {
      name: "fleet",
      logo: "rust.svg",
    },
  ];

  const swiftPM = [
    {
      name: "swift",
      logo: "swift.svg",
    },
  ];

  switch (lang) {
    case "Node.js":
    case "TypeScript":
      return nodePMs;

    case "C":
      return cPM;

    case "C++":
      return cppPM;

    case "Crystal":
      return crytsalPM;

    case "C#":
      return csharpPM;

    case "Dart":
      return dartPM;

    case "Deno":
      return denoPM;

    case "Go":
      return goPM;

    case "Java":
    case "Kotlin":
      return gradlePM;

    case "Nim":
      return nimPM;

    case "PHP":
      return phpPM;

    case "Python":
      return pythonPMs;

    case "Python":
      return pythonPMs;

    case "Ruby":
      return rubyPM;

    case "Rust":
      return rustPMs;

    case "Swift":
      return swiftPM;

    default:
      return [{ name: "default", logo: "icon.svg" }];
  }
};
