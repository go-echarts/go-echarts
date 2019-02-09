/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

const React = require('react');

const CompLibrary = require('../../core/CompLibrary.js');

const Container = CompLibrary.Container;
const GridBlock = CompLibrary.GridBlock;

class HomeSplash extends React.Component {
  render() {
    const {siteConfig, language = ''} = this.props;
    const {baseUrl, docsUrl} = siteConfig;
    const docsPart = `${docsUrl ? `${docsUrl}/` : ''}`;
    const langPart = `${language ? `${language}/` : ''}`;
    const docUrl = doc => `${baseUrl}${docsPart}${langPart}${doc}`;

    const SplashContainer = props => (
      <div className="homeContainer">
        <div className="homeSplashFade">
          <div className="wrapper homeWrapper">{props.children}</div>
        </div>
      </div>
    );

    const Logo = props => (
      <div className="projectLogo">
        <img src={props.img_src} alt="Project Logo"/>
      </div>
    );

    const ProjectTitle = () => (
      <h2 className="projectTitle">
        {siteConfig.title}
        <small>{siteConfig.tagline}</small>
      </h2>
    );

    const PromoSection = props => (
      <div className="section promoSection">
        <div className="promoRow">
          <div className="pluginRowBlock">{props.children}</div>
        </div>
      </div>
    );

    const Button = props => (
      <div className="pluginWrapper buttonWrapper">
        <a className="button" href={props.href} target={props.target}>
          {props.children}
        </a>
      </div>
    );

    return (
      <SplashContainer>
        <Logo img_src={`${baseUrl}img/logo.png`}/>
        <div className="inner">
          <ProjectTitle siteConfig={siteConfig} />
          <PromoSection>
            <Button href="docs/overview">Get start</Button>
          </PromoSection>
        </div>
      </SplashContainer>
    );
  }
}

class Index extends React.Component {
  render() {
    const {config: siteConfig, language = ''} = this.props;
    const {baseUrl} = siteConfig;

    const Block = props => (
      <Container
        padding={['bottom', 'top']}
        id={props.id}
        background={props.background}>
        <GridBlock
          align="center"
          contents={props.children}
          layout={props.layout}
        />
      </Container>
    );

    const TryOut = () => (
      <Block id="try" background="light">
        {[
          {
            content: '启动 HTTP 服务器预览所有示例图表',
            image: `${baseUrl}img/run_example.png`,
            imageAlign: 'left',
            title: '5 分钟总览',
          },
        ]}
      </Block>
    );

    const LearnHow = () => (
      <Block>
        {[
          {
            content: '只需要 5 行核心代码即可绘制出优雅的图表',
            image: `${baseUrl}img/quickstart.png`,
            imageAlign: 'right',
            title: '1 分钟上手',
          },
        ]}
      </Block>
    );

    const Features = () => (
      <Block layout="fourColumn" background="light">
        {[
          {
            content: '简洁的接口设计 使用如丝滑般流畅',
            image: `${baseUrl}img/concise.png`,
            imageAlign: 'top',
            title: '简洁',
          },
          {
            content: '囊括了 20+ 种常见图表 应有尽有',
            image: `${baseUrl}img/abundant.png`,
            imageAlign: 'top',
            title: '丰富',
          },
          {
            content: '高度灵活的配置项 可轻松搭配出精美的图表',
            image: `${baseUrl}img/flexible.png`,
            imageAlign: 'top',
            title: '灵活',
          },
        ]}
      </Block>
    );

    return (
      <div>
        <HomeSplash siteConfig={siteConfig} language={language} />
        <div className="mainContainer">
          <Features />
          <LearnHow />
          <TryOut />
        </div>
      </div>
    );
  }
}

module.exports = Index;
