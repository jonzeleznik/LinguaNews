/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './internal/view/**/*.templ',
  ],
  theme: {
    extend: {
      colors: {
        'light': {
          'bg': '#F6F1F1',
          'primary': '#AFD3E2',
          'secondary': '#19A7CE',
          'accent': '#146C94',
        },
        'dark': {
          'bg': '#2D3250',
          'primary': '#424769',
          'secondary': '#7077A1',
          'accent': '#F6B17A',
        }
      },
      fontFamily: {
        mono: ['Courier Prime', 'monospace'],
      },
    },
  },
  plugins: [],
}

