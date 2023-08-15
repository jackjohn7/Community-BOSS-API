import React from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import HomepageFeatures from '@site/src/components/HomepageFeatures';

import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <h1 className="hero__title">{siteConfig.title}</h1>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/v1/intro">
            Get Started
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={`Docs - Community BOSS API`}
      description="Community BOSS API documentation">
      <HomepageHeader />
      <main>
        <div className={styles.bodyText}>
          <h2>Documentation</h2>
          <p>We've written the documentation using Docusaurus.</p>
          <p>This project is still under active development. Beware of bugs, and please report it on the issues page of our GitHub repo or in our Discord.</p>
          <h2>What is this?</h2>
          <p>
            This project is dedicated to helping Louisiana Tech students make programming projects using data from BOSS. Before this project, there was no way
            of easily attaining data for data-driven projects relating to classes being offered at Louisiana Tech.
          </p>
          <p>Note: This is not officially endorsed by Louisiana Tech University.</p>
          <h2>Creation</h2>
          <p>
            I, Jack Branch, headed the development of this API and site. However, this project is only one part of a larger project. This API was created to
            give students access to the data that's used primarily by another application we built. This project is very much a part of the other.
          </p>
          <p>
            This project overall was lead by Cameron Thomas and myself. We've spent many months working on a web application designed to help students schedule
            for classes.
          </p>
          <p>
            We understand the struggle of scheduling classes at our university. In creating a solution to that problem, we learned the struggle of creating
            an application reliant on our school's class data. For that reason, we've created this publicly consumable API for working with data from BOSS.
          </p>
        </div>
      </main>
    </Layout>
  );
}
