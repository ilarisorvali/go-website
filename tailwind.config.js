

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/html/**/*.html",
    "./internal/content/*.go",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'), // ← correct import
  ],
}
