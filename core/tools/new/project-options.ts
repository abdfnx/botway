const languages = {
  discord: [
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
  ],
  telegram: [
    {
      name: "Node.js",
      slug: "nodejs",
    },
    {
      name: "C++",
      slug: "cpp",
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
  ],
  slack: [
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
  ],
  twitch: [
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
  ],
};

export const PLP: any = {
  discord: {
    nodejs: {
      pm: ["npm", "pnpm", "yarn"],
    },
    c: {
      pm: ["default"],
    },
    cpp: {
      pm: ["cmake"],
    },
    crystal: {
      pm: ["shards"],
    },
    csharp: {
      pm: ["dotnet"],
    },
    dart: {
      pm: ["pub"],
    },
    deno: {
      pm: ["deno package manager"],
    },
    go: {
      pm: ["go package manager"],
    },
    java: {
      pm: ["gradle"],
    },
    kotlin: {
      pm: ["gradle"],
    },
    nim: {
      pm: ["nimble"],
    },
    php: {
      pm: ["composer"],
    },
    python: {
      pm: ["pip", "pipenv", "poetry"],
    },
    ruby: {
      pm: ["bundler"],
    },
    rust: {
      pm: ["cargo"],
    },
    typescript: {
      pm: ["npm", "pnpm", "yarn"],
    },
  },
  telegram: {
    nodejs: {
      pm: ["npm", "pnpm", "yarn"],
    },
    cpp: {
      pm: ["cmake"],
    },
    csharp: {
      pm: ["dotnet"],
    },
    dart: {
      pm: ["pub"],
    },
    deno: {
      pm: ["deno package manager"],
    },
    go: {
      pm: ["go package manager"],
    },
    java: {
      pm: ["gradle"],
    },
    kotlin: {
      pm: ["gradle"],
    },
    nim: {
      pm: ["nimble"],
    },
    php: {
      pm: ["composer"],
    },
    python: {
      pm: ["pip", "pipenv", "poetry"],
    },
    ruby: {
      pm: ["bundler"],
    },
    rust: {
      pm: ["cargo"],
    },
    swift: {
      pm: ["swift"],
    },
    typescript: {
      pm: ["npm", "pnpm", "yarn"],
    },
  },
  slack: {
    nodejs: {
      pm: ["npm", "pnpm", "yarn"],
    },
    python: {
      pm: ["pip", "pipenv", "poetry"],
    },
    typescript: {
      pm: ["npm", "pnpm", "yarn"],
    },
  },
  twitch: {
    nodejs: {
      pm: ["npm", "pnpm", "yarn"],
    },

    deno: {
      pm: ["deno package manager"],
    },
    go: {
      pm: ["go package manager"],
    },
    java: {
      pm: ["gradle"],
    },
    python: {
      pm: ["pip", "pipenv", "poetry"],
    },
    typescript: {
      pm: ["npm", "pnpm", "yarn"],
    },
  },
};

export const platforms = [
  {
    name: "Choose",
    slug: "choose",
  },
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
    typeName: "Choose",
    type: "choose",
  },
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
  switch (platform) {
    case "Discord":
      return languages["discord"];

    case "Telegram":
      return languages["telegram"];

    case "Slack":
      return languages["slack"];

    case "Twitch":
      return languages["twitch"];

    default:
      return [
        {
          name: "Choose",
          slug: "choose",
        },
      ];
  }
};

export const packageManagers = (lang: any) => {
  const nodePMs = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
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
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "default",
      logo: "c.svg",
    },
  ];

  const cppPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "cmake",
      logo: "cmake.svg",
    },
  ];

  const crytsalPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "shards",
      logo: "crystal.svg",
    },
  ];

  const csharpPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "dotnet",
      logo: "dotnet.svg",
    },
  ];

  const dartPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "pub",
      logo: "dart.svg",
    },
  ];

  const denoPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "deno package manager",
      logo: "deno.svg",
    },
  ];

  const goPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "go package manager",
      logo: "go.svg",
    },
  ];

  const gradlePM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "gradle",
      logo: "gradle.svg",
    },
  ];

  const nimPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "nimble",
      logo: "nimble.svg",
    },
  ];

  const phpPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "composer",
      logo: "composer.svg",
    },
  ];

  const pythonPMs = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "pip",
      logo: "pip.svg",
    },
    {
      name: "pipenv",
      logo: "pipenv.svg",
    },
    {
      name: "poetry",
      logo: "poetry.svg",
    },
  ];

  const rubyPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "bundler",
      logo: "bundler.svg",
    },
  ];

  const rustPMs = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
    {
      name: "cargo",
      logo: "cargo.png",
    },
  ];

  const swiftPM = [
    {
      name: "Choose",
      logo: "choose.svg",
    },
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

    case "Choose":
      return [{ name: "Choose", logo: "choose.svg" }];

    default:
      return [{ name: "default", logo: "icon.svg" }];
  }
};
