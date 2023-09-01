/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./src/components/*.{js,jsx,ts,tsx}"
  ], 
  theme: {
    colors: { 
        "white": "#ffffff",
        "black": "#000000"
    },
    fontFamily: {
      'uber': ['Uber Move Text', 'sans-serif']
    },
    extend: {
      gridTemplateColumns: {
        // Simple 16 column grid
        '16': 'repeat(16, minmax(0, 1fr))',        
      },
      fontFamily: {
        'uber': ['Uber Move Text', 'sans-serif']
      }
    },
  },
  plugins: [],
}