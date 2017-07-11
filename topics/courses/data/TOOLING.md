# Go packages

All of these can be installed via `go get ...`

github.com/kniren/gota/dataframe    
go-hep.org/x/hep/csvutil   
go-hep.org/x/hep/csvutil/csvdriver   
github.com/lib/pq    
github.com/patrickmn/go-cache    
github.com/boltdb/bolt    
github.com/gonum/matrix/mat64    
github.com/gonum/stat    
github.com/montanaflynn/stats    
github.com/gonum/plot     
github.com/gonum/plot/plotter    
github.com/gonum/plot/vg     
github.com/gonum/stat/distuv    
github.com/gonum/mathext     
bitbucket.org/zombiezen/gopdf/pdf          
github.com/gonum/floats     
github.com/sajari/regression     
github.com/sjwhitworth/golearn/base     
github.com/sjwhitworth/golearn/evaluation    
github.com/sjwhitworth/golearn/knn     
github.com/sjwhitworth/golearn/trees    
github.com/Sirupsen/logrus   
github.com/satori/go.uuid     
github.com/gogo/protobuf/types     
github.com/gogo/protobuf/proto     
github.com/gogo/protobuf/gogoproto     
github.com/pachyderm/pachyderm     

# Other dependencies

- A free "tiny turtle" Postgres instance on [ElephantSQL](https://www.elephantsql.com/)
- `psql` 9.2+
- A [local installation of Pachyderm](http://pachyderm.readthedocs.io/en/latest/getting_started/local_installation.html).
- A specific version of `google.golang.com/grpc`:
    
    ```
    go get google.golang.org/grpc
    cd $GOPATH/src/google.golang.org/grpc
    git checkout 21f8ed309495401e6fd79b3a9fd549582aed1b4c 
    ```
