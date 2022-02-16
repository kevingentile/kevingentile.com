import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ArticleComponent } from './article/article.component';
import { ArticlesComponent } from './articles/articles.component';
import { HomeComponent } from './home/home.component';
import { ObsComponent } from './obs/obs.component';

const routes: Routes = [
  { path: '',
    component: HomeComponent
  },
  {
    path: 'rambler',
    component: ArticlesComponent,
  },
  {
    path: 'rambler/:articleDate',
    component: ArticleComponent
  },
  {
    path: 'obs',
    component: ObsComponent
  },
  {
    path: '**',
    redirectTo: ''
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
