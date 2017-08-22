import scss from 'rollup-plugin-scss'
import bundleSize from 'rollup-plugin-bundle-size'
import autoprefixer from 'autoprefixer'
import postcss from 'postcss'
import fs from 'fs'

const isProduction = process.env.NODE_ENV === 'production'

export default function (options) {
  return scss({
    output: function (styles) {
      postcss([autoprefixer])
        .process(styles)
        .then(result => {
          const dest = `${options.distPath}.css`
          fs.writeFileSync(dest, result.css)
          bundleSize().ongenerate({
            dest
          }, {
            code: result.css
          })
        })
    },
    verbose: true,
    includePaths: [ 'src/', 'node_modules'],
    outputStyle: isProduction ? 'compressed' : 'nested'
  })
};
