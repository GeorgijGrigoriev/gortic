const MyNews = [
  {   
      id: 1,
      author: "Test",
      text: "Test text"
  },
  {   
      id: 2,
      author: "Tester",
      text: "Test text acepted"
  }
];


<Article />

class Article extends React.Component {
    render() {
      const { author, text } = this.props.data
      return (
        <div className="row">
            <div className="col col-md-6 offset-md-3">
                <div className="article card">
                    <div className="card-body">
                        <p className="news__author card-title">{author}:</p>
                        <p className="news__text card-text">{text}</p>
                    </div>
                </div>
            </div>
        </div>
      )
    }
  }

class News extends React.Component {
    render() {
      const { data } = this.props
      let newsTemplate

      if (data.length) {
        newsTemplate = data.map(function(item) {
          return <Article key={item.id} data={item}/>
        })
      } else {
        newsTemplate = <p>К сожалению новостей нет</p>
      }
      return (
        <div className="news">
          {newsTemplate}
          {
            data.length ? <div className="col col-md-6 offset-md-3"> <strong className={'news__count'}>Всего новостей: {data.length}</strong></div> : null
          }
        </div>
      );
    }
  }


const App = () => {
      return (
          <React.Fragment>
                <News data={MyNews} />
          </React.Fragment>
  )
}
ReactDOM.render(
  <App />,
  document.getElementById('root')
);