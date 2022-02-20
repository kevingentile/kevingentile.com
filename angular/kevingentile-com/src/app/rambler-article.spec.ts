import { RamblerArticle } from './rambler-article';

describe('RamblerArticle', () => {
  it('should create an instance', () => {
    expect(new RamblerArticle("title", 0, "date", "body")).toBeTruthy();
  });
});
