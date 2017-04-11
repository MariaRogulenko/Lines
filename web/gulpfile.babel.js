"use strict"

import babelify from "babelify";
import browserify from "browserify";
import buffer from "vinyl-buffer";
import gulp from "gulp";
import notify from "gulp-notify";
import rimraf from "rimraf";
import runSequence from "run-sequence";
import source from "vinyl-source-stream";
import uglify from "gulp-uglify";
import watchify from "watchify";

const paths = {
  outputDir: "static",
  sourceJs: "app/main.js",
  outputJs: "app.js",
};

gulp.task("browserify", () => {
  // Bundles the file and all dependencies into one file.
  browserify(paths.sourceJs)
      // Makes browserify aware of the es2015 and JSX.
      .transform(babelify, {presets: ["es2015", "stage-1", "react"]})
      .bundle()
    // source() is makes browserify compatible with gulp.
    .pipe(source(paths.outputJs))
    .pipe(buffer())
    .pipe(uglify())
    .pipe(gulp.dest(paths.outputDir));
});

gulp.task("copyHtml", () => {
  gulp.src("index.html")
    .pipe(gulp.dest(paths.outputDir));
});

gulp.task("clean", cb => {
  rimraf(paths.outputDir, cb);
});

gulp.task("watchify", () => {
  const bundler = browserify(paths.sourceJs, watchify.args)
    .plugin(watchify, {ignoreWatch: ["**/node_modules/**"]});
  const rebundle = () => bundler.bundle().on("error", notify.onError())
    .pipe(source("app.js"))
    .pipe(gulp.dest(paths.outputDir));
  bundler.transform(babelify, {presets: ["es2015", "stage-1", "react"]}).on("update", rebundle);
  return rebundle();
});

gulp.task("build", cb => {
  process.env.NODE_ENV = "production";
  runSequence("clean", ["browserify", "copyHtml"], cb);
});

gulp.task("watch", cb => {
  runSequence("clean", ["watchify", "copyHtml"], cb);
});
