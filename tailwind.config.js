/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/html/**/*.html", "./internal/content/*.go"],
  theme: {
    extend: {
      fontFamily: {
        inter: ["Inter", "sans-serif"],
        merriweather: ["Merriweather", "serif"],
        monaspace: ["Monaspace", "monospace"],
      },
    },
  },
};
