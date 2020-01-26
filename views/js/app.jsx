class App extends React.Component {
    render() {
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>GovtechList</h1>
                    <p>Sign in to get access</p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            companies: []
        }
    }

    render() {
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br/>
                    <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                    <h2>GovtechList</h2>
                    <div className="row">
                        {this.state.companies.map(function(company, i){
                            return (<Company key={i} company={company}/>)
                        })}
                    </div>
                </div>
            </div>
        )
    }
}

class Company extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            liked: ""
        }
        this.like = this.like.bind(this);
    }

    like(){

    }

    render(){
        return(
            <div className="col-xs-4">
                <div className="panel panel-default">
                    <div className="panel-heading">#{this.props.company.id}<span className="pull-right">{this.state.liked}</span></div>
                    <div className="panel-body">
                        {this.props.company.company}
                    </div>
                    <div className="panel-footer">
                        {this.props.jokes.likes} Likes &nbsp;
                        <a onClick={this.like} className="btn btn-default">
                            <span className="glyphicon glyphicon-thumbs-up"></span>
                        </a>
                    </div>

                </div>
            </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));