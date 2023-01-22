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
    component: ArticleComponent,
    pathMatch: 'full'
  },
  {
    path: 'obs',
    component: ObsComponent
  },
  {
    path: '**',
    redirectTo: '' // TODO 404 component
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
