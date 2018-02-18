import { Component } from '@angular/core';
import { OnInit } from '@angular/core/src/metadata/lifecycle_hooks';
import { FormGroup, FormControl } from '@angular/forms';
import { map, catchError } from 'rxjs/operators'
import { Observable } from 'rxjs/Observable';
import { _throw } from 'rxjs/observable/throw'
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  constructor( private http: HttpClient){}

  form :FormGroup;
  headers = new HttpHeaders()
              // .set( 'Content-Type', 'application/json' )

  ngOnInit(){
    this.initForm()
  }

  
  initForm(){
    this.form = new FormGroup( { 
      'file': new FormControl( null )
    } )
  }


  onFileChange(event) {

    let headers;
    let fileList: FileList = event.target.files;
    if(fileList.length > 0) {
        let file: File = fileList[0];
        let formData:FormData = new FormData();
        formData.append('uploadFile', file, file.name);

        this.http.post(`/file`, formData, { headers } )
            .pipe( catchError( error => {
              return new Observable( obs => obs.error( error ))
            } )
            )
            .subscribe(
              data => console.log('success'),
              error => console.log(error)
            )
    }
  }
  
}
