import { Link } from "gatsby"
import PropTypes from "prop-types"
import React from "react"

const Header = ({ title, description }) => (
  <header className="logo">
          <div>
            <Link to="/">
              <h2>{title}</h2>
            </Link>
            <p>{description}</p>
            <div class="social-links">
              <a href="mailto:kwangonya@gmail.com" title="Email"><i class="fa fa-envelope"></i></a>
              <a href="https://twitter.com/wang0nya" title="Twitter" target="_blank"><i class="fab fa-twitter"></i></a>
              <a href="https://github.com/wangonya" title="Github" target="_blank"><i class="fab fa-github"></i></a>
              <a href="https://dev.to/wangonya" title="Dev" target="_blank"><i class="fab fa-dev"></i></a>
              <a href="https://www.linkedin.com/in/wangonya/" title="LinkedIn" target="_blank"><i class="fab fa-linkedin"></i></a>
              <a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ" title="Facebook" target="_blank"><i class="fab fa-facebook"></i></a>
              <a href="https://wangonya.com/rss.xml" title="RSS" target="_blank"><i class="fas fa-rss"></i></a>
            </div>
          </div>
  </header>
)

Header.propTypes = {
  avatar: PropTypes.string,
}

Header.defaultProps = {
  avatar: ``,
}

export default Header
