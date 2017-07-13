import { SPage } from './app.po';

describe('s App', () => {
  let page: SPage;

  beforeEach(() => {
    page = new SPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!');
  });
});
