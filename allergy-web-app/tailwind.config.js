/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./src/components/*.{js,jsx,ts,tsx}"
  ], 
  theme: {
    colors: { 
        "white": "#ffffff",
        "black": "#000000",
        "light-gray": "#bbbbbb", 
        "btn-gray": "#333333",
        "gray": "#eeeeee",
        "blue": "rgb(59 130 246)",
        "green": "rgb(74 222 128)",
        "red": "rgb(225 29 72)",
    },
    fontFamily: {
      'uber': ['Uber Move Text', 'sans-serif']
    },
    extend: {
      gridTemplateColumns: {
        '16': 'repeat(16, minmax(0, 1fr))',        
      },
      translate: {
        '2xfull': '200%'
      },
      gridColumnStart: {
        '13': '13',
        '14': '14',
        '15': '15',
        '16': '16',
        '17': '17',
      },
      margin: {
        '50%': '50%'
      },
      fontFamily: {
        'uber': ['Uber Move Text', 'sans-serif']
      }
    },
  },
  plugins: [],
}