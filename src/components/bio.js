/**
 * Bio component that queries for data
 * with Gatsby's useStaticQuery component
 *
 * See: https://www.gatsbyjs.org/docs/use-static-query/
 */

import React from "react"
import { useStaticQuery, graphql } from "gatsby"
import Image from "gatsby-image"

import { rhythm } from "../utils/typography"

const Bio = () => {
  const data = useStaticQuery(graphql`
    query BioQuery {
      avatar: file(absolutePath: { regex: "/me.jpg/" }) {
        childImageSharp {
          fixed(width: 50, height: 50) {
            ...GatsbyImageSharpFixed
          }
        }
      }
      site {
        siteMetadata {
          author {
            name
            summary
          }
          social {
            twitter
          }
        }
      }
    }
  `)

  const { author, social } = data.site.siteMetadata
  return (
    <div
      style={{
        display: `flex`,
        marginBottom: rhythm(2.5),
      }}
    >
      <p>
        <h3>Kinyanjui Wangonya</h3>
        {author.summary} <br/>
        <a href="mailto:kwangonya@gmail.com" title="Email">Email</a>
        {`  `}| <a href="https://github.com/wangonya" title="Github" target="_blank">Github</a>
        {`  `}| <a href="https://dev.to/wangonya" title="Dev" target="_blank">Dev</a>
        {`  `}| <a href="https://www.linkedin.com/in/wangonya/" title="LinkedIn" target="_blank">Linkedin</a>
      </p>
    </div>
  )
}

export default Bio
