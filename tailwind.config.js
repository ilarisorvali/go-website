/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/html/**/*.html", "./internal/content/*.go"],
  theme: {
    extend: {
      fontFamily: {
        roboto: ['"Roboto"', "sans-serif"],
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
