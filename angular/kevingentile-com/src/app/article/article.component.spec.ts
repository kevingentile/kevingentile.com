import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { of } from 'rxjs';
import { AppRoutingModule } from '../app-routing.module';
import { ArticleService } from '../article.service';

import { ArticleComponent } from './article.component';

describe('ArticleComponent', () => {
  let component: ArticleComponent;
  let fixture: ComponentFixture<ArticleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ArticleComponent ],
      imports: [AppRoutingModule, HttpClientTestingModule]
    })
    .compileComponents();
  });

  let articleService: ArticleService
  beforeEach(() => {
    fixture = TestBed.createComponent(ArticleComponent);
    component = fixture.componentInstance;
    articleService = TestBed.inject(ArticleService)
    fixture.detectChanges();
  });

  it('should create', () => {
    spyOn(articleService, "getArticles").and.returnValue(of([]))
    expect(component).toBeTruthy();
  });
});
