import uglify from 'rollup-plugin-uglify'

const isProduction = process.env.NODE_ENV === 'production'

export default function (options = {}) {
  let uglifyOptions = {
    output: {
      comments: function (node, comment) {
        var text = comment.value
        var type = comment.type
        if (type === 'comment2') {
          // multiline comment
          return /@preserve|@license|@cc_on/i.test(text)
        }
      }
    }
  }

  uglifyOptions.sourceMap = isProduction
    ? {}
    : {
      filename: `${options.distPath}/app.js`,
      url: 'inline'
    }

  uglifyOptions.warnings = !isProduction
  uglifyOptions.compress = isProduction
  uglifyOptions.mangle = isProduction

  return uglify(uglifyOptions)
}
