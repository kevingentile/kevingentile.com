import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ArticleComponent } from './article/article.component';
import { ArticlesComponent } from './articles/articles.component';
import { HomeComponent } from './home/home.component';

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
    path: '**',
    redirectTo: ''
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
