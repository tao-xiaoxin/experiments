# Vue-cli 与Vite 环境搭建与项目构建

在之前的语法演示中，我们直接使用 script 引入 Vue 3，从而在浏览器里实现了所有调试功能。但是在实际的项目中，我们会使用专门的调试工具。在项目上线之前，代码也需要打包压缩，并且考虑到研发效率和代码可维护性，所以在下面，我们需要建立一个工程化的项目来实现这些功能。

在早期的 Vue 2中，官方推荐用 Vue-cli 创建项目；近两年 Vue.js 官方推出了 Vite，很多新项目也开始使用 Vite。而对于 Vue 3，**官方推荐使用 Vite 创建项目**，因为 vite 能够提供更好更快的调试体验。**此外，Vue-cli 将于2023年底停止维护**。但是，有些老项目还是用的Vue-cli，我们还是有必要了解一下。

![image-20230823220212265](https://billy.taoxiaoxin.club/md/2023/08/64eaf6fa8503edde53e345ee.png)

在之前我们已经安装过 Node.js ，接下来我们开始搭建Vue CLI的环境。

## 一.Vue-cli 

### 01.基本介绍

官网地址：https://cli.vuejs.org/zh/

Vue CLI 是一个基于 Vue.js 进行快速开发的完整系统，提供：

1. 通过 `@vue/cli` 实现的交互式的项目脚手架。

2. 通过 `@vue/cli` + `@vue/cli-service-global` 实现的零配置原型开发。

3. 一个运行时依赖 (@vue/cli-service)，该依赖：

   3.1 可升级；

   3.2 基于 webpack 构建，并带有合理的默认配置；

   3.3 可以通过项目内的配置文件进行配置；

   3.4 可以通过插件进行扩展。

4. 一个丰富的官方插件集合，集成了前端生态中最好的工具。

5. 一套完全图形化的创建和管理 Vue.js 项目的用户界面。

Vue CLI 致力于将 Vue 生态中的工具基础标准化。它确保了各种构建工具能够基于智能的默认配置即可平稳衔接，这样你可以专注在撰写应用上，而不必花好几天去纠结配置的问题。与此同时，它也为每个工具提供了调整配置的灵活性，无需 eject。

### 02.安装

执行如下命令：

```bash
npm install -g @vue/cli
```

查看版本，验证是否安装成功：

```bash
vue -V
```

### 03.创建项目

执行如下命令：

```bash
vue create myapp01
```

执行后会出现如下界面：

![image-20230824184722558](https://billy.taoxiaoxin.club/md/2023/08/64eaf703ccd77cdfaad4dc26.png)

这里我们选择第三个。

![image-20230824184805057](https://billy.taoxiaoxin.club/md/2023/08/64eaf70ce83f163033a2f136.png)

然后出现如下界面：

![image-20230824185041354](https://billy.taoxiaoxin.club/md/2023/08/64eaf714bfa82583bd7f35b0.png)

解释如下：

当您使用 Vue CLI 创建项目时，您会被要求选择预设选项和项目功能。以下是您所提到的每个选项的解释：

1. **Babel**：Babel 是一个 JavaScript 编译器，用于将新版本的 JavaScript 代码转换为旧版本的代码，以便在更旧的浏览器中运行。选择此选项会为您的项目配置 Babel。
2. **TypeScript**：TypeScript 是一个类型安全的 JavaScript 超集，它添加了静态类型检查等功能。选择此选项会为您的项目配置 TypeScript 而不仅仅是普通的 JavaScript。
3. **Progressive Web App (PWA) Support**：选择此选项可以为您的应用添加渐进式 Web 应用（PWA）的支持。PWA 允许您创建能够像原生应用一样运行的 Web 应用，包括脱机访问、通知等功能。
4. **Router**：Vue Router 是 Vue.js 官方的路由管理器，允许您为应用添加前端路由，从而创建单页面应用（SPA）。
5. **Vuex**：Vuex 是 Vue.js 官方的状态管理库，用于管理应用的状态，包括状态的存储、更新和同步。
6. **CSS Pre-processors**：选择此选项可以为您的项目集成 CSS 预处理器，如Sass、Less或Stylus，以提供更强大的 CSS 编写体验。
7. **Linter / Formatter**：选择此选项会为您的项目配置代码检查工具（Linter）和代码格式化工具。这有助于确保代码的一致性和质量。
8. **Unit Testing**：选择此选项可以为您的项目集成单元测试框架，如Jest或Mocha，以便在开发过程中对代码进行自动化测试。
9. **E2E Testing**：选择此选项可以为您的项目集成端到端（E2E）测试框架，如Cypress或Nightwatch，用于模拟用户在应用中的实际操作。

这里我们可以使用**空格键来选中或者取消**，最后按Enter键。

![image-20230824185803025](https://billy.taoxiaoxin.club/md/2023/08/64eaf71ac440e8a5f2d44864.png)

然后选择Vue 版本，这里肯定是选择Vue 3 版本，因为Vue 2 即将停止维护了。

![image-20230824190016255](https://billy.taoxiaoxin.club/md/2023/08/64eaf71f6d655383d6c263fd.png)

![image-20230824191107354](https://billy.taoxiaoxin.club/md/2023/08/64eaf728cd985609ada6bc3b.png)

继续又出现了如下界面：

![image-20230824191152553](https://billy.taoxiaoxin.club/md/2023/08/64eaf73182f4f32b9c202f30.png)

选项解释如下：

- **ESLint with error prevention only**: 这将配置 ESLint 以仅阻止潜在的代码错误，但不会对代码进行格式化。
- **ESLint + Airbnb config**: 这将配置 ESLint 并使用 Airbnb 的代码风格规范，帮助您遵循一致的代码编写方式。
- **ESLint + Standard config**: 这将配置 ESLint 并使用 Standard 的代码风格规范，类似于 Airbnb，也是一种常用的规范。
- **ESLint + Prettier**: 这将配置 ESLint 并集成 Prettier，Prettier 是一个代码格式化工具，可以自动格式化代码以遵循统一的代码样式。

这里我们选择默认就好，继续又出现了如下界面：

![image-20230824191619727](https://billy.taoxiaoxin.club/md/2023/08/64eaf7d60dfc7d6da1ba340e.png)

解释如下：

- **Lint on save**： 选择此选项后，项目会在您保存文件时自动进行代码检查。这可以帮助您在编辑代码时及时发现可能的问题，并确保代码符合您所配置的规范和风格。
- **Lint and fix on commit**： 选择此选项后，项目会在您提交代码时自动进行代码检查，并尝试自动修复一些可以自动修复的问题。例如，如果有一些格式问题，工具会尝试自动将其修复。这有助于确保您提交的代码在提交之前已经经过检查和修复。

这里我们选择默认就好，然后又出现了如下界面，选择放置配置文件：

![image-20230824192053358](https://billy.taoxiaoxin.club/md/2023/08/64eaf8bb41e87bf255d735ec.png)

解释如下：

- **In dedicated config files**： 选择此选项时，每个工具（如 Babel、ESLint 等）的配置将会被放置在项目中的独立配置文件中。例如，您会在项目根目录下找到一个 `.babelrc` 文件，以及一个 `.eslintrc` 文件，这些文件将包含相应工具的配置选项。这使得配置更加模块化和清晰，便于管理和维护。
- **In package.json**： 选择此选项时，工具的配置将会直接放置在项目的 `package.json` 文件中，通过特定的字段来定义。例如，您可能会在 `package.json` 文件中找到一个名为 `babel` 的字段，其中包含 Babel 的配置信息。这种方法减少了额外的配置文件，但可能会使 `package.json` 文件变得较大，同时也可能会降低配置的可读性。

继续选择默认，走你。继续如下界面：

![image-20230824192432215](https://billy.taoxiaoxin.club/md/2023/08/64eaf99b909db6270fad5947.png)

这里的意思是是否想将当前的项目配置保存为一个预设，以便在将来的项目中重复使用。预设是一组预定义的项目配置，在这里我们不需要，所以输入N即可。

接下来就是等待项目创建完成。

![image-20230824205206886](https://billy.taoxiaoxin.club/md/2023/08/64eaf744cced63ebc6001276.png)

看到这，就是项目创建成功了

![image-20230824205315767](https://billy.taoxiaoxin.club/md/2023/08/64eaf74eae0ef7b80aca8d10.png)

运行如下命令，启动项目：

```bash
cd myapp01
npm run serve
```

点击链接地址

![image-20230824205529348](https://billy.taoxiaoxin.club/md/2023/08/64eaf754a9397d0f874e7e44.png)

在浏览器打开后如下：

![image-20230824205618454](https://billy.taoxiaoxin.club/md/2023/08/64eaf758e50c26f3f53ff75b.png)

到此，项目创建就成功了。

### 04.目录结构

生成的项目目录结构如下：

```bash
myapp01/            // 项目根目录
├── node_modules/   // 存放项目依赖的 Node.js 模块，这些模块在 package.json 中定义。
├── public/         // 公共静态文件目录，包含主 HTML 文件和其他静态资源。
│   ├── index.html  // 主 HTML 文件，是应用的入口点。
│   └── ...         // 可以放置其他静态资源，如图像、字体等。
├── src/            // 源代码目录，包含您编写的实际应用代码。
│   ├── assets/     // 静态资源目录，存放应用使用的图像、样式文件等。
│   ├── components/ // Vue 组件目录，存放可复用的 Vue 组件。
│   │   └── ...     // 每个组件通常由 .vue 文件组成。
│   ├── views/      // 页面组件目录，存放页面级别的 Vue 组件。
│   │   └── ...     // 每个视图对应一个 .vue 文件，通常用于路由页面。
│   ├── App.vue     // 应用的根组件，定义应用的整体结构和布局。
│   └── main.js     // 应用的入口文件，初始化 Vue 实例并挂载根组件。
├── .gitignore      // Git 忽略文件列表，定义哪些文件和目录不应被版本控制。
├── babel.config.js // Babel 配置文件，定义将新版 JavaScript 转换为向后兼容的代码。
├── package.json    // 项目配置文件，包含项目信息、依赖、脚本等。
├── README.md       // 项目说明文件，通常包含项目介绍、使用方法等信息。
├── vue.config.js   // 可选的自定义 Vue CLI 配置文件，用于配置构建选项、代理等。
├── .eslintrc.js    // 可选的 ESLint 规则配置文件，定义代码检查工具的规则。
└── .prettierrc     // 可选的 Prettier 格式化配置文件，用于代码格式保持一致。
```

这些目录和文件的解释如下：

- `node_modules/`：存放项目依赖的 Node.js 模块，这些模块在 `package.json` 中定义。
- `public/`：包含一些公共静态文件，其中最重要的是 `index.html`，作为应用的主 HTML 文件。这里的文件不会经过 webpack 打包，而是直接复制到构建输出目录。
- `src/`：源代码目录，是您实际编写应用的地方。
  - `assets/`：存放静态资源，如图片、样式文件等。
  - `components/`：存放可复用的 Vue 组件，这些组件可以在应用的多个地方使用。
  - `views/`：存放页面级别的 Vue 组件，通常对应应用的不同路由页面。
  - `App.vue`：应用的根组件，它包含了应用的整体布局和结构。
  - `main.js`：应用的入口文件，初始化 Vue 实例并挂载根组件。
- `.gitignore`：Git 忽略文件列表，定义了哪些文件或目录不应该被 Git 版本控制。
- `babel.config.js`：Babel 配置文件，定义了如何将新版本的 JavaScript 转换为向后兼容的代码。
- `package.json`：项目的配置文件，包含项目信息、依赖、脚本等。
- `README.md`：项目的说明文件，通常包含项目介绍、使用方法等信息。
- `vue.config.js`：可选的自定义 Vue CLI 配置文件，用于配置构建选项、代理设置、自定义 webpack 配置等。
- `.eslintrc.js`：可选的 ESLint 规则配置文件，定义代码检查工具的规则，以确保代码遵循一致的编码风格和最佳实践。
- `.prettierrc`：可选的 Prettier 格式化配置文件，定义 Prettier 的格式化配置，用于保持代码的一致格式。

### 05.package.json 文件解析

```bash
{
  "name": "myapp01",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "serve": "vue-cli-service serve",
    "build": "vue-cli-service build",
    "lint": "vue-cli-service lint"
  },
  "dependencies": {
    "core-js": "^3.8.3",
    "vue": "^3.2.13"
  },
  "devDependencies": {
    "@babel/core": "^7.12.16",
    "@babel/eslint-parser": "^7.12.16",
    "@vue/cli-plugin-babel": "~5.0.0",
    "@vue/cli-plugin-eslint": "~5.0.0",
    "@vue/cli-service": "~5.0.0",
    "eslint": "^7.32.0",
    "eslint-plugin-vue": "^8.0.3"
  }
}
```

- `"name": "myapp01"`：这是应用程序的名称，被设置为 "myapp01"。

- `"version": "0.1.0"`：这是应用程序的版本号，被设置为 "0.1.0"。

- `"private": true`：这个字段指示该应用程序是一个私有项目，这意味着它不能被发布到公共的包管理器（如 npm）中。

- `"scripts"`：这是一个对象，定义了可以通过命令行运行的各种脚本命令。

  - `"serve"`：运行 Vue CLI 的服务命令，用于在开发模式下启动一个本地开发服务器。
  - `"build"`：运行 Vue CLI 的构建命令，用于将应用程序打包成生产环境可用的文件。
  - `"lint"`：运行 Vue CLI 的 lint 命令，用于静态代码分析和检查代码风格。

- `"dependencies"`：这是一个对象，列出了项目的生产环境依赖项（在应用程序运行时需要的库）。

  - `"core-js": "^3.8.3"`：用于提供 JavaScript 的新特性和方法的库，版本号至少为 3.8.3。
  - `"vue": "^3.2.13"`：Vue.js 框架的核心库，版本号至少为 3.2.13。

- `"devDependencies"`：这是一个对象，列出了项目的开发环境依赖项（在开发和构建过程中需要的库）。

  - `"@babel/core": "^7.12.16"`：Babel 编译工具的核心库，版本号至少为 7.12.16。
  - `"@babel/eslint-parser": "^7.12.16"`：用于在 ESLint 中解析 Babel 语法的库，版本号至少为 7.12.16。
  - `"@vue/cli-plugin-babel": "~5.0.0"`：Vue CLI 的 Babel 插件，版本号大致在 5.0.0 左右。
  - `"@vue/cli-plugin-eslint": "~5.0.0"`：Vue CLI 的 ESLint 插件，版本号大致在 5.0.0 左右。
  - `"@vue/cli-service": "~5.0.0"`：Vue CLI 的核心服务插件，版本号大致在 5.0.0 左右。
  - `"eslint": "^7.32.0"`：ESLint 静态代码分析工具，版本号至少为 7.32.0。
  - `"eslint-plugin-vue": "^8.0.3"`：用于在 ESLint 中检查 Vue.js 代码的插件，版本号至少为 8.0.3。

  

## 二.Vite

### 01.基本介绍

官网:https://cn.vitejs.dev

#### 1.1 现实问题

  在浏览器支持 ES 模块之前，JavaScript 并没有提供原生机制让开发者以模块化的方式进行开发。这也正是我们对 “打包” 这个概念熟悉的原因：使用工具抓取、处理并将我们的源码模块串联成可以在浏览器中运行的文件。

  时过境迁，我们见证了诸如 [webpack](https://webpack.js.org/)、[Rollup](https://rollupjs.org/) 和 [Parcel](https://parceljs.org/) 等工具的变迁，它们极大地改善了前端开发者的开发体验。

  然而，当我们开始构建越来越大型的应用时，需要处理的 JavaScript 代码量也呈指数级增长。包含数千个模块的大型项目相当普遍。基于 JavaScript 开发的工具就会开始遇到性能瓶颈：通常需要很长时间（甚至是几分钟！）才能启动开发服务器，即使使用模块热替换（HMR），文件修改后的效果也需要几秒钟才能在浏览器中反映出来。如此循环往复，迟钝的反馈会极大地影响开发者的开发效率和幸福感。

  **Vite 旨在利用生态系统中的新进展解决上述问题：浏览器开始原生支持 ES 模块，且越来越多 JavaScript 工具使用编译型语言编写。**

  #### 1.2 缓慢的服务器启动

  当冷启动开发服务器时，基于打包器的方式启动必须优先抓取并构建你的整个应用，然后才能提供服务。

  Vite 通过在一开始将应用中的模块区分为 **依赖** 和 **源码** 两类，改进了开发服务器启动时间。

  - **依赖** 大多为在开发时不会变动的纯 JavaScript。一些较大的依赖（例如有上百个模块的组件库）处理的代价也很高。依赖也通常会存在多种模块化格式（例如 ESM 或者 CommonJS）。

    Vite 将会使用 [esbuild](https://esbuild.github.io/) [预构建依赖](https://cn.vitejs.dev/guide/dep-pre-bundling.html)。esbuild 使用 Go 编写，并且比以 JavaScript 编写的打包器预构建依赖快 10-100 倍。

  - **源码** 通常包含一些并非直接是 JavaScript 的文件，需要转换（例如 JSX，CSS 或者 Vue/Svelte 组件），时常会被编辑。同时，并不是所有的源码都需要同时被加载（例如基于路由拆分的代码模块）。

    Vite 以 [原生 ESM](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Modules) 方式提供源码。这实际上是让浏览器接管了打包程序的部分工作：Vite 只需要在浏览器请求源码时进行转换并按需提供源码。根据情景动态导入代码，即只在当前屏幕上实际使用时才会被处理。

  ![image-20230827115357744](https://billy.taoxiaoxin.club/md/2023/08/64eaf766c5441931c654c07d.png)

  ![image-20230827115449051](https://billy.taoxiaoxin.club/md/2023/08/64eaf76aafe112597ca04182.png)

#### 1.3 缓慢的更新

  基于打包器启动时，重建整个包的效率很低。原因显而易见：因为这样更新速度会随着应用体积增长而直线下降。

  一些打包器的开发服务器将构建内容存入内存，这样它们只需要在文件更改时使模块图的一部分失活[[1\]](https://cn.vitejs.dev/guide/why.html#footnote-1)，但它也仍需要整个重新构建并重载页面。这样代价很高，并且重新加载页面会消除应用的当前状态，所以打包器支持了动态模块热替换（HMR）：允许一个模块 “热替换” 它自己，而不会影响页面其余部分。这大大改进了开发体验 —— 然而，在实践中我们发现，即使采用了 HMR 模式，其热更新速度也会随着应用规模的增长而显著下降。

  在 Vite 中，HMR 是在原生 ESM 上执行的。当编辑一个文件时，Vite 只需要精确地使已编辑的模块与其最近的 HMR 边界之间的链失活[[1\]](https://cn.vitejs.dev/guide/why.html#footnote-1)（大多数时候只是模块本身），使得无论应用大小如何，HMR 始终能保持快速更新。

  Vite 同时利用 HTTP 头来加速整个页面的重新加载（再次让浏览器为我们做更多事情）：源码模块的请求会根据 `304 Not Modified` 进行协商缓存，而依赖模块请求则会通过 `Cache-Control: max-age=31536000,immutable` 进行强缓存，因此一旦被缓存它们将不需要再次请求。

  一旦你体验到 Vite 的神速，你是否愿意再忍受像曾经那样使用打包器开发就要打上一个大大的问号了。

#### 1.4 为什么生产环境仍需打包

  尽管原生 ESM 现在得到了广泛支持，但由于嵌套导入会导致额外的网络往返，在生产环境中发布未打包的 ESM 仍然效率低下（即使使用 HTTP/2）。为了在生产环境中获得最佳的加载性能，最好还是将代码进行 tree-shaking、懒加载和 chunk 分割（以获得更好的缓存）。

  要确保开发服务器和生产环境构建之间的最优输出和行为一致并不容易。所以 Vite 附带了一套 [构建优化](https://cn.vitejs.dev/guide/features.html#build-optimizations) 的 [构建命令](https://cn.vitejs.dev/guide/build.html)，开箱即用。

#### 1.5 为何不用 ESBuild 打包？

  Vite 目前的插件 API 与使用 `esbuild` 作为打包器并不兼容。尽管 `esbuild` 速度更快，但 Vite 采用了 Rollup 灵活的插件 API 和基础建设，这对 Vite 在生态中的成功起到了重要作用。目前来看，我们认为 Rollup 提供了更好的性能与灵活性方面的权衡。

  即便如此，`esbuild` 在过去几年有了很大进展，我们不排除在未来使用 `esbuild` 进行生产构建的可能性。我们将继续利用他们所发布的新功能，就像我们在 JS 和 CSS 最小化压缩方面所做的那样，`esbuild` 使 Vite 在避免对其生态造成干扰的同时获得了性能提升。

### 02.兼容性问题

Vite 需要 [Node.js](https://nodejs.org/en/) 版本 14.18+，16+。然而，有些模板需要依赖更高的 Node 版本才能正常运行，当你的包管理器发出警告时，请注意升级你的 Node 版本。

### 03.创建项目

使用 NPM:

```bash
npm create vite@latest
```

使用 Yarn:

```bash
yarn create vite
```

使用 PNPM:

```bash
pnpm create vite
```

第一次使用会提示你下载Vite，输入Y即可。

![image-20230827121142073](https://billy.taoxiaoxin.club/md/2023/08/64eaf77216795e0ac3a6583d.png)

然后会提示你输入项目名称

![image-20230827121339761](https://billy.taoxiaoxin.club/md/2023/08/64eaf776870509b6877fecd4.png)

然后选择模板，选择Vue 模板即可

![image-20230827121540505](https://billy.taoxiaoxin.club/md/2023/08/64eaf778551e4a0008515094.png)

然后选择语言，JavaScript即可

![image-20230827121730429](https://billy.taoxiaoxin.club/md/2023/08/64eaf77efd5654fecdc32d34.png)

最后这个项目就创建成功了,运行如下命令即可:

```bash
  cd myapp02
  npm install
  npm run dev
```

点击链接地址

![image-20230827122410802](https://billy.taoxiaoxin.club/md/2023/08/64eaf786d8ece2d17ded4cae.png)

在浏览器中打开页面如下:

![image-20230827122447243](https://billy.taoxiaoxin.club/md/2023/08/64eaf7883044100d5760d1c8.png)

### 04.目录结构

```tree
myapp02/
├── README.md               # 项目的说明文档
├── index.html              # 项目的主 HTML 文件，Vite 插入构建的脚本和样式
├── package-lock.json       # 锁定依赖版本的文件
├── package.json            # 项目的配置文件，包含项目元数据、依赖和脚本等信息
├── public/                 # 不需要构建处理的静态文件
│   └── vite.svg            # 示例静态文件
├── src/                    # 项目的源代码文件
│   ├── App.vue             # 项目的根组件
│   ├── assets/             # 存放静态资源
│   │   └── vue.svg         # 示例图片资源
│   ├── components/         # 存放 Vue 组件文件
│   │   └── HelloWorld.vue # 示例 Vue 组件
│   ├── main.js             # 项目的入口文件，初始化 Vue 应用程序
│   └── style.css           # 示例样式文件
├── node_modules/           # 包含项目依赖模块
└── vite.config.js          # Vite 的配置文件

```

以下是详细解释说明:
- myapp02/项目的根目录。
  - `README.md`：项目的说明文档。
  - `index.html`：项目的主 HTML 文件，Vite 会在构建中插入脚本和样式。
  - `package-lock.json`：锁定依赖版本的文件，确保一致性。
  - `package.json`：项目的配置文件，包含项目的元数据、依赖和脚本。
  
- public/：不需要构建处理的静态文件。
  - `vite.svg`：示例静态文件，可以直接在 HTML 中引用。
  
- src/：项目的源代码文件。
  - `App.vue`：项目的根组件。
  
  - assets/：存放静态资源，如图片、样式。
    - `vue.svg`：示例图片资源。
    
  - components/：存放 Vue 组件文件。

    - `HelloWorld.vue`：示例 Vue 组件。

  - `main.js`：项目的入口文件，初始化 Vue 应用程序。

  - `style.css`：示例样式文件。

- `node_modules/`：包含项目的依赖模块，通过 `npm i` 安装。

- `vite.config.js`：Vite 的配置文件，用于自定义 Vite 的行为和设置。

## 三.Webpack 和 Vite 有什么区别？

我在研究技术的时候经常在想，脱离技术的定位来对比技术好坏都是耍流氓。因为每一种流行的技术之所以被人接纳，**肯定是有其诞生的定位和开发者的使用定位。**

所以我们要对比 Webpack 和 Vite，最重要是对比这两种技术工具的定位。Webpack 和 Vite 的定位是不一样的，这个连 Vite 的作者尤雨溪老师都曾经在知乎上回应过。**Vite 定位是 Web“开发工具链”，其内置了一些打包构建工具，让开发者开箱即用**，例如预设了 Web 开发模式直接使用 ESM 能力，开发过程中可以通过浏览器的 ESM 能力按需加载当前开发页面的相关资源。

然而，**Webpack 定位是构建“打包工具”**，面向的是前端代码的编译打包过程。Webpack 能力很单一，就是提供一个打包构建的能力，如果有特定的构建需要，必须让开发者来选择合适的 Loader 和 Plugin 进行组合配置，达到最终的想要的打包效果。比如 Webpack 没有内置开发服务，需要引入 webpack-dev-server 才能有开发服务的能力，这个对比 Vite 就不一样，VIte 就直接内置了一个开发服务。

那么，两者的技术能力或者功能有什么异同点呢？

其实这两个工具能提供的技术能力有很大的重叠度或者相似度，基本就是对前端代码进行打包构建处理。区别是 **Vite 内置了很多工具，可以减少很多配置工作量**；而 **Webpack 只是简单的打包工具架子**，需要开发者一开始准备很多配置处理，不像 Vite 那样能开箱即用，需要花些功夫进行选择 Webpack 的 Loader 和 Plugin 进行配置。

## 四.Vite 的优势

Vite 是 Vue.js 作者尤雨溪早年基于 Rollup 做的一个开发工具，核心是为了提高 JavaScript 项目的开发效率。那么相比同类型的开发工具来说，除了Vite 支持开箱即用，无需像 Webpack 要做一堆繁杂的配置之外，很重要的一点就是 Vite 确实能够提升我们的开发效率。

Vite 利用现在最新版本的浏览器支持 ESM 的特性，可以在开发模式下直接让所有 npm 模块和项目里的 JavaScript 文件按需加载执行，减少了开发模式编译时间，让开发过程中每次修改代码后能快速编译，进而提升了开发效率。

一点对大部分开发者来讲，都是解决了开发过程中很大的体验痛点。那我们再深一步，为什么 Vite 能在开发模式中快速编译代码呢？因为 Vite 用了 esbuild。而 esbuild 是用 Go 语言编写的构建器，和普通 JavaScript 实现的构建器相比，它的打包速度能快 10~100 倍。
