import uglify from 'rollup-plugin-uglify'

const isProduction = process.env.NODE_ENV === 'production'

export default function (options = {}) {
  let uglifyOptions = {}

  uglifyOptions.sourceMap = isProduction ? {} : {
    filename: `${options.distPath}/bundle.js`,
    url: 'inline'
  }

  uglifyOptions.warnings = !isProduction
  uglifyOptions.compress = isProduction
  uglifyOptions.mangle = isProduction

  return uglify(uglifyOptions)
}
