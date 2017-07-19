import { NgModule, } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { HttpModule } from '@angular/http';

import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
// import { DBModule } from '@ngrx/db';
import { RouterStoreModule } from '@ngrx/router-store';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
// import { MaterialModule } from '@angular/material';

import { ProbeService } from './services/probe.service'
import { D3Service } from './services/d3.service'

import { AppComponent } from './app.component';
// import { ComponentsModule } from './components';
// import { BookEffects } from './effects/book';
// import { CollectionEffects } from './effects/collection';
// import { BookExistsGuard } from './guards/book-exists';

import { routes } from './routes';
import { reducer } from './reducers';
import { ProbeComponent } from './components/probe/probe.component';
import { GraphComponent } from './components/graph/graph.component';
// import { schema } from './db';

@NgModule({
  declarations: [
    AppComponent,
    ProbeComponent,
    GraphComponent
  ],
  imports: [
    CommonModule,
    BrowserModule,
    HttpModule,
    RouterModule.forRoot(routes, { useHash: true }),
    StoreModule.provideStore(reducer),
    RouterStoreModule.connectRouter(),
    StoreDevtoolsModule.instrumentOnlyWithExtension()                 
  ],
  providers: [
    D3Service,
    ProbeService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
