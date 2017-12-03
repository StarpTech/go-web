import json from 'rollup-plugin-json'
import resolve from 'rollup-plugin-node-resolve'
import babel from 'rollup-plugin-babel'
import commonjs from 'rollup-plugin-commonjs'
import postcss from 'postcss'
import riot from 'rollup-plugin-riot'
import uglify from './rollup.uglify.config'
import replace from 'rollup-plugin-replace'
import postcssCssnext from 'postcss-cssnext'
import scssTask from './rollup.scss.config'

const appDist = 'dist/app'

const config = {
  input: 'src/app.js',
  plugins: [
    resolve({ jsnext: true }),
    scssTask({
      distPath: appDist
    }),
    json(),
    riot({
      style: 'cssnext',
      type: 'es6',
      parsers: {
        css: { cssnext }
      }
    }),
    replace({
      exclude: 'node_modules/**',
      ENV: JSON.stringify(process.env.NODE_ENV || 'development')
    }),
    commonjs(),
    babel(),
    // uglify()
  ],
  output: {
    file: `${appDist}.js`,
    format: 'iife',
    name: 'App',
    sourcemap: true
  },
  watch: {
    exclude: ['node_modules/**']
  }
}

/**
 * Transforms new CSS specs into more compatible CSS
 */
function cssnext (tagName, css) {
  // A small hack: it passes :scope as :root to PostCSS.
  // This make it easy to use css variables inside tags.
  css = css.replace(/:scope/g, ':root')
  css = postcss([postcssCssnext]).process(css).css
  css = css.replace(/:root/g, ':scope')
  return css
}

export default [config]
