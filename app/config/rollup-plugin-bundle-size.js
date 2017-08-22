const path = require('path')
const chalk = require('chalk')
const maxmin = require('maxmin')

export default function (options) {
  return {
    name: 'rollup-plugin-bundle-size',
    ongenerate (details, result) {
      const asset = path.basename(details.file)
      const size = maxmin(result.code, result.code, true)
      console.log(`Created bundle ${chalk.cyan(asset)}: ${size.substr(size.indexOf(' → ') + 3)}`)
    }
  }
}
