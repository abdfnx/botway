module.exports = {
  mode: "jit",
  content: ["./**/*.{js,jsx,ts,tsx}"],
  darkMode: "class",
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [
    require("tailwindcss-animate"),
  ],
};
