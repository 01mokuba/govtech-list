// import Companies from './components/companies';

class App extends React.Component {
  state = {
    companies: []
  }

  constructor() {
    super();
    fetch("http://localhost:3000/api/companies")
        .then(res => res.json())
        .then(companies => this.setState({companies}))
  }

  render() {
      const companies = this.state.companies;
      return (
          <div className="container">
              <div className="col-lg-12">
                  <br/>
                  <h2>GovtechList</h2>
                  <div className="row row-eq-height">
                      <div className="panel-body">
                          {companies.map(company =>
                              
                                  <div className="panel panel-default">
                                      <div className="panel-heading"><a href={company.official_url}>{company.company}</a>（{company.founded_year}）</div>
                                      <div className="panel-body">{company.description}</div>
                                  </div>
                              
                          )}
                      </div>
                  </div>
              </div>
          </div>
      );
  }
}

ReactDOM.render(<App />, document.getElementById('app'));