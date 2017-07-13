import { Routes } from '@angular/router';

import { AppComponent } from './app.component';
// import { NotFoundPageComponent } from './containers/not-found-page';

export const routes: Routes = [
  {
    path: '',
    component: AppComponent
  }
  // {
  //   path: '**',
  //   component: NotFoundPageComponent
  // }
];